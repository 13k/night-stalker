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

func (app *App) serveLiveMatches(c echo.Context) error {
	cc := c.(*nswebctx.Context)

	view, err := app.loadLiveMatchesView()

	if err != nil {
		app.log.WithError(err).Error("error loading LiveMatches view")

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

func (app *App) loadLiveMatchesView() (*nspb.LiveMatches, error) {
	matchIDs, err := app.rds.LiveMatchIDs()

	if err != nil {
		return nil, xerrors.Errorf("error fetching live match IDs from redis: %w", err)
	}

	data, err := app.dbl.LiveMatchesData(app.ctx, &nsdbda.LiveMatchesParams{
		MatchIDs:         matchIDs,
		WithFollowedOnly: true,
	})

	if err != nil {
		return nil, xerrors.Errorf("error loading live matches data: %w", err)
	}

	view, err := nsvs.NewLiveMatches(data)

	if err != nil {
		return nil, xerrors.Errorf("error creating live matches view: %w", err)
	}

	return view, nil
}
