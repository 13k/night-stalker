package web

import (
	"errors"
	"net/http"

	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	nsbus "github.com/13k/night-stalker/internal/bus"
	nspb "github.com/13k/night-stalker/internal/protocol"
	nsrds "github.com/13k/night-stalker/internal/redis"
	nsviews "github.com/13k/night-stalker/internal/views"
	"github.com/13k/night-stalker/models"
)

var (
	errRedisPubSubSubscription = errors.New("Redis pubsub subscription failed")
)

func (app *App) subscribeLiveMatchesUpdate() error {
	app.rdsSubLiveMatchesUpdate = app.rds.Subscribe(nsrds.TopicLiveMatchesUpdate)

	msg, err := app.rdsSubLiveMatchesUpdate.Receive()

	if err != nil {
		return err
	}

	if _, ok := msg.(*redis.Subscription); !ok {
		return errRedisPubSubSubscription
	}

	go app.watchLiveMatchesUpdate()

	return nil
}

func (app *App) watchLiveMatchesUpdate() {
	defer func() {
		app.rdsSubLiveMatchesUpdate.Close()
		app.log.Warn("stopped watching live matches")
	}()

	for {
		select {
		case <-app.ctx.Done():
			return
		case msg, ok := <-app.rdsSubLiveMatchesUpdate.Channel():
			if !ok {
				return
			}

			go app.handleLiveMatchesUpdate(msg)
		}
	}
}

func (app *App) handleLiveMatchesUpdate(rmsg *redis.Message) {
	l := app.log.WithFields(logrus.Fields{
		"channel": rmsg.Channel,
		"pattern": rmsg.Pattern,
		"payload": rmsg.Payload,
	})

	l.Debug("received live matches update")

	matchIDs, err := app.liveMatchIDs()

	if err != nil {
		l.WithError(err).Error("error loading live matches ids")
		return
	}

	view, err := app.loadLiveMatchesView(matchIDs...)

	if err != nil {
		l.WithError(err).Error("error loading live matches view")
		return
	}

	if len(view.Matches) == 0 {
		l.Debug("ignoring empty live matches update")
		return
	}

	busmsg := &nsbus.LiveMatchesChangeMessage{
		Change: &nspb.LiveMatchesChange{
			Op:     nspb.LiveMatchesChange_REPLACE,
			Change: view,
		},
	}

	app.bus.Pub(busmsg, nsbus.TopicLiveMatches)
}

func (app *App) liveMatchIDs() ([]nspb.MatchID, error) {
	idx, err := app.rds.Get(nsrds.KeyLiveMatchesIndex).Int()

	if err != nil {
		app.log.
			WithField("key", nsrds.KeyLiveMatchesIndex).
			WithError(err).
			Error("error fetching cached live matches index")

		return nil, err
	}

	liveMatchesKey := nsrds.KeyLiveMatches(idx)
	result := app.rds.ZRevRange(liveMatchesKey, 0, -1)

	if err := result.Err(); err != nil {
		app.log.
			WithField("key", liveMatchesKey).
			WithError(err).
			Error("error fetching cached live matches index")

		return nil, err
	}

	matchIDs := make([]nspb.MatchID, len(result.Val()))

	if err := result.ScanSlice(&matchIDs); err != nil {
		app.log.
			WithError(err).
			Error("error parsing live match IDs")

		return nil, err
	}

	return matchIDs, nil
}

func (app *App) loadLiveMatchesView(matchIDs ...nspb.MatchID) (*nspb.LiveMatches, error) {
	if len(matchIDs) == 0 {
		return &nspb.LiveMatches{}, nil
	}

	matches, err := app.loadLiveMatches(matchIDs)

	if err != nil {
		return nil, err
	}

	if len(matches) == 0 {
		return &nspb.LiveMatches{}, nil
	}

	matchIDs = make([]nspb.MatchID, len(matches))
	accountIDs := make([]nspb.AccountID, 0, len(matches)*10)

	for i, match := range matches {
		matchIDs[i] = match.MatchID

		for _, player := range match.Players {
			accountIDs = append(accountIDs, player.AccountID)
		}
	}

	stats, err := app.loadLiveMatchStats(matchIDs)

	if err != nil {
		return nil, err
	}

	statsByMatchID := make(map[nspb.MatchID]*models.LiveMatchStats, len(stats))

	for _, stat := range stats {
		statsByMatchID[stat.MatchID] = stat
	}

	followedPlayers, err := app.loadFollowedPlayers(accountIDs)

	if err != nil {
		return nil, err
	}

	accountIDs = make([]nspb.AccountID, 0, len(followedPlayers))

	for accountID := range followedPlayers {
		accountIDs = append(accountIDs, accountID)
	}

	players, err := app.loadPlayers(accountIDs)

	if err != nil {
		return nil, err
	}

	proPlayers, err := app.loadProPlayers(accountIDs)

	if err != nil {
		return nil, err
	}

	view, err := nsviews.NewLiveMatches(
		matches,
		statsByMatchID,
		followedPlayers,
		players,
		proPlayers,
	)

	if err != nil {
		return nil, err
	}

	return view, nil
}

func (app *App) loadLiveMatches(matchIDs []nspb.MatchID) ([]*models.LiveMatch, error) {
	var matches []*models.LiveMatch

	err := app.db.
		Joins("INNER JOIN live_match_players ON live_matches.id = live_match_players.live_match_id").
		Joins("INNER JOIN followed_players ON live_match_players.account_id = followed_players.account_id").
		Where("live_matches.match_id IN (?)", matchIDs).
		Group("live_matches.id").
		Order("live_matches.sort_score DESC").
		Preload("Players").
		Find(&matches).
		Error

	if err != nil {
		app.log.WithError(err).Error("database live matches")
		return nil, err
	}

	return matches, nil
}

func (app *App) loadLiveMatchStats(matchIDs []nspb.MatchID) ([]*models.LiveMatchStats, error) {
	var stats []*models.LiveMatchStats

	subQuery := app.db.
		Model(models.LiveMatchStatsModel).
		Select("id, row_number() OVER(PARTITION BY match_id ORDER BY created_at DESC) AS row_id").
		Where("match_id IN (?)", matchIDs).
		SubQuery()

	err := app.db.
		Joins("INNER JOIN (?) latest ON (live_match_stats.id = latest.id and latest.row_id = ?)", subQuery, 1).
		Preload("Teams").
		Preload("Players").
		Find(&stats).
		Error

	if err != nil {
		app.log.WithError(err).Error("database stats")
		return nil, err
	}

	return stats, nil
}

func (app *App) loadFollowedPlayers(accountIDs []nspb.AccountID) (map[nspb.AccountID]*models.FollowedPlayer, error) {
	var players []*models.FollowedPlayer

	err := app.db.
		Where("account_id IN (?)", accountIDs).
		Find(&players).
		Error

	if err != nil {
		app.log.WithError(err).Error("database followed players")
		return nil, err
	}

	byAccountID := make(map[nspb.AccountID]*models.FollowedPlayer, len(players))

	for _, player := range players {
		byAccountID[player.AccountID] = player
	}

	return byAccountID, nil
}

func (app *App) loadPlayers(accountIDs []nspb.AccountID) (map[nspb.AccountID]*models.Player, error) {
	var players []*models.Player

	err := app.db.
		Where("account_id IN (?)", accountIDs).
		Find(&players).
		Error

	if err != nil {
		app.log.WithError(err).Error("database players")
		return nil, err
	}

	byAccountID := make(map[nspb.AccountID]*models.Player, len(players))

	for _, player := range players {
		byAccountID[player.AccountID] = player
	}

	return byAccountID, nil
}

func (app *App) loadProPlayers(accountIDs []nspb.AccountID) (map[nspb.AccountID]*models.ProPlayer, error) {
	var players []*models.ProPlayer

	err := app.db.
		Where("account_id IN (?)", accountIDs).
		Find(&players).
		Error

	if err != nil {
		app.log.WithError(err).Error("database pro players")
		return nil, err
	}

	byAccountID := make(map[nspb.AccountID]*models.ProPlayer, len(players))

	for _, player := range players {
		byAccountID[player.AccountID] = player
	}

	return byAccountID, nil
}

func (app *App) serveLiveMatches(c echo.Context) error {
	matchIDs, err := app.liveMatchIDs()

	if err != nil {
		return err
	}

	view, err := app.loadLiveMatchesView(matchIDs...)

	if err != nil {
		app.log.WithError(err).Error("error loading live matches view")

		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		}
	}

	return c.JSON(http.StatusOK, view)
}
