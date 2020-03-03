package web

import (
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsdb "github.com/13k/night-stalker/internal/db"
	nsviews "github.com/13k/night-stalker/internal/views"
	"github.com/13k/night-stalker/models"
)

func (app *App) loadHeroesData(heroIDs ...models.HeroID) ([]*models.Hero, error) {
	var heroes []*models.Hero

	scope := app.db.Debug()

	if len(heroIDs) > 0 {
		scope = scope.Where("id IN (?)", heroIDs)
	}

	err := scope.Find(&heroes).Error

	if err != nil {
		err = xerrors.Errorf("error loading heroes: %w", err)
		return nil, err
	}

	return heroes, nil
}

func (app *App) loadHeroMatchesData(id models.HeroID) (nsviews.MatchesData, error) {
	matchIDs, err := nsdb.FindMatchIDs(app.db, nsdb.FindMatchIDsFilters{
		PlayerFilters: &nsdb.PlayerFilters{
			HeroIDs: nscol.HeroIDs{id},
		},
	})

	if err != nil {
		err = xerrors.Errorf("error finding match IDs: %w", err)
		return nil, err
	}

	if len(matchIDs) == 0 {
		return nil, nil
	}

	matchesData, err := nsdb.LoadMatchesData(app.db, matchIDs...)

	if err != nil {
		err = xerrors.Errorf("error loading matches data: %w", err)
		return nil, err
	}

	return matchesData, nil
}
