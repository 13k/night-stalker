package web

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protocol"
	nsviews "github.com/13k/night-stalker/internal/views"
	"github.com/13k/night-stalker/models"
)

func (app *App) serveLiveMatches(c echo.Context) error {
	matchIDs, err := app.rdsLiveMatchIDs()

	if err != nil {
		app.log.Error(fmt.Sprintf("%+v", err))

		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		}
	}

	view, err := app.loadLiveMatchesView(matchIDs...)

	if err != nil {
		app.log.Error(fmt.Sprintf("%+v", err))

		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		}
	}

	return c.JSON(http.StatusOK, view)
}

func (app *App) createLiveMatchesView(liveMatches ...*models.LiveMatch) (*nspb.LiveMatches, error) {
	if len(liveMatches) == 0 {
		return &nspb.LiveMatches{}, nil
	}

	matchIDs := make(nscol.MatchIDs, len(liveMatches))
	accountIDs := make([]nspb.AccountID, 0, len(liveMatches)*10)

	for i, liveMatch := range liveMatches {
		matchIDs[i] = liveMatch.MatchID

		for _, player := range liveMatch.Players {
			accountIDs = append(accountIDs, player.AccountID)
		}
	}

	stats, err := app.loadLiveMatchStats(matchIDs)

	if err != nil {
		err = xerrors.Errorf("error loading live match stats: %w", err)
		return nil, err
	}

	statsByMatchID := make(map[nspb.MatchID]*models.LiveMatchStats, len(stats))

	for _, stat := range stats {
		statsByMatchID[stat.MatchID] = stat
	}

	followedPlayers, err := app.loadFollowedPlayers(accountIDs)

	if err != nil {
		err = xerrors.Errorf("error loading followed players: %w", err)
		return nil, err
	}

	accountIDs = make([]nspb.AccountID, 0, len(followedPlayers))

	for accountID := range followedPlayers {
		accountIDs = append(accountIDs, accountID)
	}

	players, err := app.loadPlayers(accountIDs)

	if err != nil {
		err = xerrors.Errorf("error loading players: %w", err)
		return nil, err
	}

	proPlayers, err := app.loadProPlayers(accountIDs)

	if err != nil {
		err = xerrors.Errorf("error loading pro players: %w", err)
		return nil, err
	}

	view, err := nsviews.NewLiveMatches(
		liveMatches,
		statsByMatchID,
		followedPlayers,
		players,
		proPlayers,
	)

	if err != nil {
		err = xerrors.Errorf("error creating LiveMatches view: %w", err)
		return nil, err
	}

	return view, nil
}

func (app *App) loadLiveMatchesView(matchIDs ...nspb.MatchID) (*nspb.LiveMatches, error) {
	if len(matchIDs) == 0 {
		return &nspb.LiveMatches{}, nil
	}

	liveMatches, err := app.loadLiveMatches(matchIDs)

	if err != nil {
		err = xerrors.Errorf("error loading live matches: %w", err)
		return nil, err
	}

	view, err := app.createLiveMatchesView(liveMatches...)

	if err != nil {
		err = xerrors.Errorf("error creating LiveMatches view: %w", err)
		return nil, err
	}

	return view, nil
}

func (app *App) loadLiveMatches(matchIDs nscol.MatchIDs) ([]*models.LiveMatch, error) {
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
		err = xerrors.Errorf("error loading matches: %w", err)
		return nil, err
	}

	return matches, nil
}

func (app *App) loadLiveMatchStats(matchIDs nscol.MatchIDs) ([]*models.LiveMatchStats, error) {
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
		err = xerrors.Errorf("error loading live match stats: %w", err)
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
		err = xerrors.Errorf("error loading followed players: %w", err)
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
		err = xerrors.Errorf("error loading players: %w", err)
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
		err = xerrors.Errorf("error loading pro players: %w", err)
		return nil, err
	}

	byAccountID := make(map[nspb.AccountID]*models.ProPlayer, len(players))

	for _, player := range players {
		byAccountID[player.AccountID] = player
	}

	return byAccountID, nil
}
