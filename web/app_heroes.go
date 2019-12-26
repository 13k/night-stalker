package web

import (
	"net/http"

	"github.com/labstack/echo/v4"

	nspb "github.com/13k/night-stalker/internal/protocol"
	nsviews "github.com/13k/night-stalker/internal/views"
	"github.com/13k/night-stalker/models"
)

func (app *App) loadHeroesView() ([]*nspb.Hero, error) {
	var heroes []*models.Hero

	err := app.db.Find(&heroes).Error

	if err != nil {
		app.log.WithError(err).Error("database heroes")
		return nil, err
	}

	view := nsviews.NewHeroes(heroes)

	return view, nil
}

func (app *App) serveHeroes(c echo.Context) error {
	heroes, err := app.loadHeroesView()

	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		}
	}

	return c.JSON(http.StatusOK, heroes)
}
