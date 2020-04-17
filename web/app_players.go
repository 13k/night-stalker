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

func (app *App) servePlayerMatches(c echo.Context) error {
	cc := c.(*nswebctx.Context)

	type PathParams struct {
		AccountID nspb.AccountID `param:"account_id"`
	}

	pathParams := &PathParams{}

	if err := cc.Bind(pathParams); err != nil {
		return &echo.HTTPError{
			Code:     http.StatusBadRequest,
			Message:  "invalid account id",
			Internal: err,
		}
	}

	view, err := app.loadPlayerMatchesView((*nsdbda.PlayerMatchesParams)(pathParams))

	if err != nil {
		app.log.WithError(err).Error("error loading PlayerMatches view")

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

func (app *App) loadPlayerMatchesView(params *nsdbda.PlayerMatchesParams) (*nspb.PlayerMatches, error) {
	data, err := app.dbl.PlayerMatchesData(app.ctx, params)

	if err != nil {
		return nil, xerrors.Errorf("error loading PlayerMatches data: %w", err)
	}

	view, err := nsvs.NewPlayerMatches(data)

	if err != nil {
		return nil, xerrors.Errorf("error creating PlayerMatches view: %w", err)
	}

	return view, nil
}
