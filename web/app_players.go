package web

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	nspb "github.com/13k/night-stalker/internal/protocol"
	nsviews "github.com/13k/night-stalker/internal/views"
	"github.com/13k/night-stalker/models"
)

func (app *App) loadPlayerView(accountID nspb.AccountID) (*nspb.Player, error) {
	followed := &models.FollowedPlayer{}

	err := app.db.
		Where(&models.FollowedPlayer{AccountID: accountID}).
		Take(followed).
		Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		app.log.WithError(err).Error("database followed player")
		return nil, err
	}

	player := &models.Player{}

	err = app.db.
		Where(&models.Player{AccountID: accountID}).
		Take(player).
		Error

	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			app.log.WithError(err).Error("database player")
			return nil, err
		}

		player = nil
	}

	proPlayer := &models.ProPlayer{}

	err = app.db.
		Where(&models.ProPlayer{AccountID: accountID}).
		Preload("Team").
		Take(proPlayer).
		Error

	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			app.log.WithError(err).Error("database pro player")
			return nil, err
		}

		proPlayer = nil
	}

	type livePlayerResult struct {
		*models.LiveMatchPlayer
		MatchID nspb.MatchID
	}

	livePlayerResults := []*livePlayerResult{}

	err = app.db.
		Table(models.LiveMatchPlayerModel.TableName()).
		Select("live_match_players.*, live_matches.match_id").
		Joins("INNER JOIN live_matches ON (live_match_players.live_match_id = live_matches.id)").
		Where(&models.LiveMatchPlayer{AccountID: accountID}).
		Scan(&livePlayerResults).
		Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		app.log.WithError(err).Error("database live players")
		return nil, err
	}

	livePlayers := make(map[nspb.MatchID]*models.LiveMatchPlayer)

	for _, result := range livePlayerResults {
		livePlayers[result.MatchID] = result.LiveMatchPlayer
	}

	type statsPlayerResult struct {
		*models.LiveMatchStatsPlayer
		MatchID nspb.MatchID
	}

	statsPlayerResults := []*statsPlayerResult{}

	err = app.db.
		Table(models.LiveMatchStatsPlayerModel.TableName()).
		Select("live_match_stats_players.*, live_match_stats.match_id").
		Joins("INNER JOIN live_match_stats ON (live_match_stats_players.live_match_stats_id = live_match_stats.id)").
		Where(&models.LiveMatchStatsPlayer{AccountID: accountID}).
		Scan(&statsPlayerResults).
		Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		app.log.WithError(err).Error("database stats players")
		return nil, err
	}

	statsPlayers := make(map[nspb.MatchID]*models.LiveMatchStatsPlayer)

	for _, result := range statsPlayerResults {
		statsPlayers[result.MatchID] = result.LiveMatchStatsPlayer
	}

	view := nsviews.NewPlayer(
		followed,
		player,
		proPlayer,
		livePlayers,
		statsPlayers,
	)

	return view, nil
}

func (app *App) servePlayer(c echo.Context) error {
	type PathParams struct {
		AccountID nspb.AccountID `param:"account_id"`
	}

	pathParams := &PathParams{}

	if err := c.Bind(pathParams); err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  "invalid account id",
			Internal: err,
		}
	}

	view, err := app.loadPlayerView(pathParams.AccountID)

	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		}
	}

	if view == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, view)
}
