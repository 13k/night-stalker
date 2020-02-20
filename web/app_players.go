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

	var matchPlayers []*models.MatchPlayer

	err = app.db.
		Where(&models.MatchPlayer{AccountID: accountID}).
		Preload("Match").
		Find(&matchPlayers).
		Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		app.log.WithError(err).Error("database players")
		return nil, err
	}

	var livePlayers []*models.LiveMatchPlayer

	err = app.db.
		Where(&models.LiveMatchPlayer{AccountID: accountID}).
		Preload("LiveMatch").
		Find(&livePlayers).
		Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		app.log.WithError(err).Error("database live players")
		return nil, err
	}

	livePlayersByMatchID := make(map[nspb.MatchID]*models.LiveMatchPlayer)

	for _, livePlayer := range livePlayers {
		livePlayersByMatchID[livePlayer.LiveMatch.MatchID] = livePlayer
	}

	var statsPlayers []*models.LiveMatchStatsPlayer

	err = app.db.
		Where(&models.LiveMatchStatsPlayer{AccountID: accountID}).
		Preload("LiveMatchStats").
		Find(&statsPlayers).
		Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		app.log.WithError(err).Error("database stats players")
		return nil, err
	}

	statsPlayersByMatchID := make(map[nspb.MatchID][]*models.LiveMatchStatsPlayer)

	for _, statsPlayer := range statsPlayers {
		statsPlayersByMatchID[statsPlayer.LiveMatchStats.MatchID] =
			append(statsPlayersByMatchID[statsPlayer.LiveMatchStats.MatchID], statsPlayer)
	}

	return nsviews.NewPlayer(
		followed,
		player,
		proPlayer,
		matchPlayers,
		livePlayersByMatchID,
		statsPlayersByMatchID,
	)
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
