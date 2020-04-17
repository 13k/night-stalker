package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsdbda "github.com/13k/night-stalker/internal/db/dataaccess"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsvs "github.com/13k/night-stalker/internal/views"
	nswebctx "github.com/13k/night-stalker/web/internal/context"
)

func (app *App) serveLeagues(c echo.Context) error {
	cc := c.(*nswebctx.Context)

	params := &struct {
		LeagueIDs nscol.LeagueIDs `query:"id"`
	}{}

	if err := cc.Bind(params); err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  "invalid league id(s)",
			Internal: err,
		}
	}

	view, err := app.loadLeaguesView(params.LeagueIDs)

	if err != nil {
		app.log.WithError(err).Error("error loading Leagues view")

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

func (app *App) loadLeaguesView(leagueIDs nscol.LeagueIDs) (*nspb.Leagues, error) {
	leagues, err := app.dbl.Leagues(app.ctx, nsdbda.LeagueFilters{
		LeagueIDs: leagueIDs,
	})

	if err != nil {
		return nil, xerrors.Errorf("error loading leagues data: %w", err)
	}

	if len(leagues) == 0 {
		return nil, nil
	}

	view, err := nsvs.NewLeagues(leagues)

	if err != nil {
		return nil, xerrors.Errorf("error creating leagues views: %w", err)
	}

	return view, nil
}
