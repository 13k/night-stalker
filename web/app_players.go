package web

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protocol"
	nsviews "github.com/13k/night-stalker/internal/views"
)

func (app *App) servePlayerMatches(c echo.Context) error {
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

	view, err := app.loadPlayerMatchesView(pathParams.AccountID)

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

func (app *App) loadPlayerMatchesView(accountID nspb.AccountID) (*nspb.PlayerMatches, error) {
	playersData, err := app.loadPlayersData(accountID)

	if err != nil {
		err = xerrors.Errorf("error loading players data: %w", err)
		return nil, err
	}

	if playersData == nil {
		return nil, nil
	}

	playerData := playersData[accountID]

	if playerData == nil {
		return nil, nil
	}

	matchesData, err := app.loadPlayerMatchesData(accountID)

	if err != nil {
		err = xerrors.Errorf("error loading player matches data: %w", err)
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

	view, err := nsviews.NewPlayerMatches(
		playerData,
		knownPlayersData,
		matchesData,
	)

	if err != nil {
		err = xerrors.Errorf("error creating PlayerMatches view: %w", err)
		return nil, err
	}

	return view, nil
}
