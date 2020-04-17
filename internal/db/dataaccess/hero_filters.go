package dataaccess

import (
	"github.com/doug-martin/goqu/v9"
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsdb "github.com/13k/night-stalker/internal/db"
	nsm "github.com/13k/night-stalker/models"
)

var _ nsdb.SelectQueryFilter = (*HeroFilters)(nil)

var (
	EmptyHeroFilters = HeroFilters{}
)

type HeroFilters struct {
	HeroIDs nscol.HeroIDs
	Query   string
}

func (f HeroFilters) Empty() bool {
	return len(f.HeroIDs) == 0 && f.Query == ""
}

func (f HeroFilters) Validate() error {
	if f.Empty() {
		return xerrors.Errorf("invalid hero filters: %w", ErrEmptyFilters)
	}

	return nil
}

func (f HeroFilters) Filter(q *nsdb.SelectQuery) *nsdb.SelectQuery {
	if l := len(f.HeroIDs); l > 0 {
		pk := nsm.HeroTable.PK()

		if l == 1 {
			q = q.Eq(pk, f.HeroIDs[0])
		} else {
			q = q.In(pk, f.HeroIDs)
		}

		q = q.Prepared(true)
	}

	if f.Query != "" {
		likePattern := "%" + f.Query + "%"

		q = q.
			Where(goqu.Or(
				nsm.HeroTable.Col("localized_name").ILike(likePattern),
				goqu.Any(f.Query).Eq(nsm.HeroTable.Col("aliases")),
			)).
			Prepared(true)
	}

	return q
}
