package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"

	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsvs "github.com/13k/night-stalker/internal/views"
	nswebctx "github.com/13k/night-stalker/web/internal/context"
)

func (app *App) serveHeroes(c echo.Context) error {
	cc := c.(*nswebctx.Context)

	view, err := app.loadHeroesView()

	if err != nil {
		app.log.WithError(err).Error("error loading Heroes view")

		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		}
	}

	if view == nil {
		return cc.NoContent(http.StatusNoContent)
	}

	return cc.RespondWith(http.StatusOK, view)
}

func (app *App) serveHeroMatches(c echo.Context) error {
	cc := c.(*nswebctx.Context)

	type PathParams struct {
		HeroID nspb.HeroID `param:"id"`
	}

	pathParams := &PathParams{}

	if err := cc.Bind(pathParams); err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  "invalid hero id",
			Internal: err,
		}
	}

	view, err := app.loadHeroMatchesView((*nsdbda.HeroMatchesParams)(pathParams))

	if err != nil {
		app.log.WithError(err).Error("error loading HeroMatches view")

		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		}
	}

	if view == nil {
		return cc.NoContent(http.StatusNotFound)
	}

	return cc.RespondWith(http.StatusOK, view)
}

func (app *App) loadHeroesView() (*nspb.Heroes, error) {
	heroes, err := app.dbl.Heroes(app.ctx, nsdbda.EmptyHeroFilters)

	if err != nil {
		return nil, xerrors.Errorf("error loading heroes data: %w", err)
	}

	if len(heroes) == 0 {
		return nil, nil
	}

	view := nsvs.NewHeroes(heroes)

	return view, nil
}

func (app *App) loadHeroMatchesView(params *nsdbda.HeroMatchesParams) (*nspb.HeroMatches, error) {
	data, err := app.dbl.HeroMatchesData(app.ctx, params)

	if err != nil {
		return nil, xerrors.Errorf("error loading HeroMatches data: %w", err)
	}

	view, err := nsvs.NewHeroMatches(data)

	if err != nil {
		return nil, xerrors.Errorf("error creating HeroMatches view: %w", err)
	}

	return view, nil
}
