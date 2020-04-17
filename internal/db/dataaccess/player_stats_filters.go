package dataaccess

import (
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsdb "github.com/13k/night-stalker/internal/db"
	nsm "github.com/13k/night-stalker/models"
)

// PlayerStatsFilters holds filters for `MatchPlayer`, `LiveMatchPlayer` and `LiveMatchStatsPlayer`.
//
// It cannot be used directly as a `QueryFilterer` but can be used by one.
type PlayerStatsFilters struct {
	AccountIDs nscol.AccountIDs
	HeroIDs    nscol.HeroIDs
}

func (f PlayerStatsFilters) Empty() bool {
	return len(f.AccountIDs) == 0 && len(f.HeroIDs) == 0
}

func (f PlayerStatsFilters) Validate() error {
	if f.Empty() {
		return xerrors.Errorf("invalid player filters: %w", ErrEmptyFilters)
	}

	return nil
}

func (f PlayerStatsFilters) Filter(q *nsdb.SelectQuery, m nsm.Model) *nsdb.SelectQuery {
	t := m.Table()

	if l := len(f.AccountIDs); l > 0 {
		colAccountID := t.Col("account_id")

		if l == 1 {
			q = q.Eq(colAccountID, f.AccountIDs[0])
		} else {
			q = q.In(colAccountID, f.AccountIDs)
		}

		q = q.Prepared(true)
	}

	if l := len(f.HeroIDs); l > 0 {
		colHeroID := t.Col("hero_id")

		if l == 1 {
			q = q.Eq(colHeroID, f.HeroIDs[0])
		} else {
			q = q.In(colHeroID, f.HeroIDs)
		}

		q = q.Prepared(true)
	}

	return q
}
