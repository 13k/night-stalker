package web

import (
	"net/http"

	"github.com/labstack/echo/v4"

	nspb "github.com/13k/night-stalker/internal/protocol"
	nsviews "github.com/13k/night-stalker/internal/views"
	"github.com/13k/night-stalker/models"
)

func (app *App) loadSearchView(query string) (*nspb.Search, error) {
	var heroes []*models.Hero

	likePattern := "%" + query + "%"

	err := app.db.
		Where("localized_name ILIKE ?", likePattern).
		Find(&heroes).
		Error

	if err != nil {
		app.log.WithError(err).Error("database heroes")
		return nil, err
	}

	var followed []*models.FollowedPlayer

	err = app.db.
		Where("label ILIKE ?", likePattern).
		Find(&followed).
		Error

	if err != nil {
		app.log.WithError(err).Error("database followed players")
		return nil, err
	}

	accountIDs := make([]nspb.AccountID, len(followed))

	for i, fp := range followed {
		accountIDs[i] = fp.AccountID
	}

	var players []*models.Player

	err = app.db.
		Where("account_id IN (?)", accountIDs).
		Find(&players).
		Error

	if err != nil {
		app.log.WithError(err).Error("database players")
		return nil, err
	}

	var proPlayers []*models.ProPlayer

	err = app.db.
		Where("account_id IN (?)", accountIDs).
		Find(&proPlayers).
		Error

	if err != nil {
		app.log.WithError(err).Error("database pro players")
		return nil, err
	}

	return nsviews.NewSearch(heroes, followed, players, proPlayers)
}

func (app *App) serveSearch(c echo.Context) error {
	view, err := app.loadSearchView(c.QueryParam("q"))

	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		}
	}

	return c.JSON(http.StatusOK, view)
}
