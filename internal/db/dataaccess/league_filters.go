package dataaccess

import (
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsdb "github.com/13k/night-stalker/internal/db"
	nsm "github.com/13k/night-stalker/models"
)

var _ nsdb.SelectQueryFilter = (*LeagueFilters)(nil)

type LeagueFilters struct {
	LeagueIDs nscol.LeagueIDs
	Query     string
}

func (f LeagueFilters) Empty() bool {
	return len(f.LeagueIDs) == 0 && f.Query == ""
}

func (f LeagueFilters) Validate() error {
	if f.Empty() {
		return xerrors.Errorf("invalid league filters: %w", ErrEmptyFilters)
	}

	return nil
}

func (f LeagueFilters) Filter(q *nsdb.SelectQuery) *nsdb.SelectQuery {
	if l := len(f.LeagueIDs); l > 0 {
		pk := nsm.LeagueTable.PK()

		if l == 1 {
			q = q.Eq(pk, f.LeagueIDs[0])
		} else {
			q = q.In(pk, f.LeagueIDs)
		}

		q = q.Prepared(true)
	}

	if f.Query != "" {
		likePattern := "%" + f.Query + "%"
		q = q.
			ILike(nsm.LeagueTable.Col("name"), likePattern).
			Prepared(true)
	}

	return q
}
