package dataaccess

import (
	"github.com/doug-martin/goqu/v9/exp"
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsdb "github.com/13k/night-stalker/internal/db"
	nsm "github.com/13k/night-stalker/models"
)

var _ nsdb.SelectQueryFilter = (*LiveMatchFilters)(nil)

type LiveMatchFilters struct {
	MatchIDs         nscol.MatchIDs
	WithFollowedOnly bool
	OrderBy          nsdb.OrderByFields
}

func (f LiveMatchFilters) Empty() bool {
	return len(f.MatchIDs) == 0
}

func (f LiveMatchFilters) Validate() error {
	if f.Empty() {
		return xerrors.Errorf("invalid live match filters: %w", ErrEmptyFilters)
	}

	return nil
}

func (f LiveMatchFilters) Filter(q *nsdb.SelectQuery) *nsdb.SelectQuery {
	tLiveMatch := nsm.LiveMatchTable
	tLiveMatchPlayer := nsm.LiveMatchPlayerTable
	tFollowedPlayer := nsm.FollowedPlayerTable

	if l := len(f.MatchIDs); l > 0 {
		colMatchID := tLiveMatch.Col("match_id")

		if l == 1 {
			q = q.Eq(colMatchID, f.MatchIDs[0])
		} else {
			q = q.In(colMatchID, f.MatchIDs)
		}

		q = q.Prepared(true)
	}

	if f.WithFollowedOnly {
		q = q.
			InnerJoinEq(
				tLiveMatch.PK(),
				tLiveMatchPlayer.Col("live_match_id"),
			).
			InnerJoinEq(
				tLiveMatchPlayer.Col("account_id"),
				tFollowedPlayer.Col("account_id"),
			).
			GroupBy(tLiveMatch.PK())
	}

	if len(f.OrderBy) > 0 {
		order := f.OrderBy.MapOrderBy(func(f string) exp.Orderable {
			return tLiveMatch.Col(f)
		})

		q = q.Order(order...)
	}

	return q
}
