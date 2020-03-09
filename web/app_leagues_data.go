package web

import (
	"golang.org/x/xerrors"

	"github.com/13k/night-stalker/models"
)

func (app *App) loadLeaguesData(leagueIDs ...models.LeagueID) ([]*models.League, error) {
	var leagues []*models.League

	scope := app.db

	if len(leagueIDs) > 0 {
		scope = scope.Where("id IN (?)", leagueIDs)
	}

	err := scope.Find(&leagues).Error

	if err != nil {
		err = xerrors.Errorf("error loading leagues: %w", err)
		return nil, err
	}

	return leagues, nil
}
