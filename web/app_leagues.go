package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsviews "github.com/13k/night-stalker/internal/views"
	nswebctx "github.com/13k/night-stalker/web/internal/context"
)

func (app *App) serveLeagues(c echo.Context) error {
	cc := c.(*nswebctx.Context)

	params := &struct {
		LeagueIDs []nspb.LeagueID `query:"id"`
	}{}

	if err := cc.Bind(params); err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  "invalid league id(s)",
			Internal: err,
		}
	}

	view, err := app.loadLeaguesView(params.LeagueIDs...)

	if err != nil {
		app.log.WithError(err).Error("error loading Leagues view")
		app.log.Errorx(err)

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

func (app *App) loadLeaguesView(leagueIDs ...nspb.LeagueID) ([]*nspb.League, error) {
	data, err := app.loadLeaguesData(leagueIDs...)

	if err != nil {
		err = xerrors.Errorf("error loading leagues data: %w", err)
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	view, err := nsviews.NewLeagues(data)

	if err != nil {
		err = xerrors.Errorf("error creating leagues views: %w", err)
		return nil, err
	}

	return view, nil
}
