package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsviews "github.com/13k/night-stalker/internal/views"
	"github.com/13k/night-stalker/models"
	nswebctx "github.com/13k/night-stalker/web/internal/context"
)

func (app *App) serveSearch(c echo.Context) error {
	cc := c.(*nswebctx.Context)

	view, err := app.loadSearchView(c.QueryParam("q"))

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

func (app *App) loadSearchView(query string) (*nspb.Search, error) {
	var heroes []*models.Hero

	likePattern := "%" + query + "%"

	err := app.db.
		Where("localized_name ILIKE ? OR ? = ANY(aliases)", likePattern, query).
		Find(&heroes).
		Error

	if err != nil {
		err = xerrors.Errorf("error loading heroes: %w", err)
		return nil, err
	}

	var followed []*models.FollowedPlayer

	err = app.db.
		Where("label ILIKE ?", likePattern).
		Find(&followed).
		Error

	if err != nil {
		err = xerrors.Errorf("error loading followed players: %w", err)
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
		err = xerrors.Errorf("error loading players: %w", err)
		return nil, err
	}

	var proPlayers []*models.ProPlayer

	err = app.db.
		Where("account_id IN (?)", accountIDs).
		Find(&proPlayers).
		Error

	if err != nil {
		err = xerrors.Errorf("error loading pro players: %w", err)
		return nil, err
	}

	view := nsviews.NewSearch(heroes, followed, players, proPlayers)

	return view, nil
}
