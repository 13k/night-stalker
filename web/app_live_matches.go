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

func (app *App) loadLiveMatchesView() ([]*nspb.LiveMatch, error) {
	matchIDs, err := app.liveMatchIDs()

	if err != nil {
		return nil, err
	}

	matches, err := app.loadLiveMatches(matchIDs)

	if err != nil {
		return nil, err
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

func (app *App) subscribeLiveMatches() error {
	app.rdsSubLiveMatches = app.rds.Subscribe(nsrds.TopicLiveMatchesUpdate)

	msg, err := app.rdsSubLiveMatches.Receive()

	if err != nil {
		return err
	}

	if _, ok := msg.(*redis.Subscription); !ok {
		return errRedisPubSubSubscription
	}

	go app.watchLiveMatches()

	return nil
}

func (app *App) watchLiveMatches() {
	defer func() {
		app.rdsSubLiveMatches.Close()
		app.log.Warn("stopped watching live matches")
	}()

	for {
		select {
		case <-app.ctx.Done():
			return
		case msg, ok := <-app.rdsSubLiveMatches.Channel():
			if !ok {
				return
			}

			go app.handleLiveMatches(msg)
		}
	}
}

func (app *App) handleLiveMatches(rmsg *redis.Message) {
	app.log.WithFields(logrus.Fields{
		"channel": rmsg.Channel,
		"pattern": rmsg.Pattern,
		"payload": rmsg.Payload,
	}).Debug("received live matches")

	matches, err := app.loadLiveMatchesView()

	if err != nil {
		app.log.WithError(err).Error("error loading live matches")
		return
	}

	busmsg := &nsbus.LiveMatchesProtoMessage{Matches: matches}

	app.bus.Pub(busmsg, nsbus.TopicLiveMatches)
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

/*
	select s.*
	from live_match_stats s
	inner join (
		select id, row_number() over(partition by match_id order by created_at desc) as row_id
		from live_match_stats
		where match_id in (5159282511, 5159280849)
	) latest on (s.id = latest.id and latest.row_id = 1);
*/
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
	matches, err := app.loadLiveMatchesView()

	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		}
	}

	return c.JSON(http.StatusOK, matches)
}
