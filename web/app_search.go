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

func (app *App) serveSearch(c echo.Context) error {
	cc := c.(*nswebctx.Context)

	view, err := app.loadSearchView(&nsdbda.SearchParams{
		Query: c.QueryParam("q"),
	})

	if err != nil {
		app.log.WithError(err).Error("error loading Search view")

		return &echo.HTTPError{
			Code:     http.StatusInternalServerError,
			Message:  err.Error(),
			Internal: err,
		}
	}

	return cc.RespondWith(http.StatusOK, view)
}

func (app *App) loadSearchView(params *nsdbda.SearchParams) (*nspb.Search, error) {
	data, err := app.dbl.SearchData(app.ctx, params)

	if err != nil {
		return nil, xerrors.Errorf("error loading Search data: %w", err)
	}

	view := nsvs.NewSearch(data)

	return view, nil
}
