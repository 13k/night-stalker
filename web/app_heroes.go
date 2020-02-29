package web

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsviews "github.com/13k/night-stalker/internal/views"
	"github.com/13k/night-stalker/models"
)

func (app *App) serveHeroes(c echo.Context) error {
	view, err := app.loadHeroesView()

	if err != nil {
		app.log.Error(fmt.Sprintf("%+v", err))

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

func (app *App) serveHeroMatches(c echo.Context) error {
	type PathParams struct {
		ID models.HeroID `param:"id"`
	}

	pathParams := &PathParams{}

	if err := c.Bind(pathParams); err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  "invalid hero id",
			Internal: err,
		}
	}

	view, err := app.loadHeroMatchesView(pathParams.ID)

	if err != nil {
		app.log.Error(fmt.Sprintf("%+v", err))

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

func (app *App) loadHeroesView() ([]*nspb.Hero, error) {
	data, err := app.loadHeroesData()

	if err != nil {
		err = xerrors.Errorf("error loading heroes data: %w", err)
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	view := nsviews.NewHeroes(data)

	return view, nil
}

func (app *App) loadHeroMatchesView(id models.HeroID) (*nspb.HeroMatches, error) {
	heroesData, err := app.loadHeroesData(id)

	if err != nil {
		err = xerrors.Errorf("error loading heroes data: %w", err)
		return nil, err
	}

	if len(heroesData) == 0 {
		return nil, nil
	}

	heroData := heroesData[0]
	matchesData, err := app.loadHeroMatchesData(id)

	if err != nil {
		err = xerrors.Errorf("error loading hero matches data: %w", err)
		return nil, err
	}

	playersAccountIDs := matchesData.AccountIDs()

	var followedAccountIDs nscol.AccountIDs

	if len(playersAccountIDs) > 0 {
		followedAccountIDs, err = app.filterFollowedPlayersAccountIDs(playersAccountIDs...)

		if err != nil {
			err = xerrors.Errorf("error filtering followed players account IDs: %w", err)
			return nil, err
		}
	}

	var knownPlayersData nsviews.PlayersData

	if len(followedAccountIDs) > 0 {
		knownPlayersData, err = app.loadPlayersData(followedAccountIDs...)

		if err != nil {
			err = xerrors.Errorf("error loading players data: %w", err)
			return nil, err
		}
	}

	view, err := nsviews.NewHeroMatches(
		heroData,
		knownPlayersData,
		matchesData,
	)

	if err != nil {
		err = xerrors.Errorf("error creating HeroMatches view: %w", err)
		return nil, err
	}

	return view, nil
}
