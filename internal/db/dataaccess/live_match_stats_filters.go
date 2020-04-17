package dataaccess

import (
	"github.com/doug-martin/goqu/v9"
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsdb "github.com/13k/night-stalker/internal/db"
	nsm "github.com/13k/night-stalker/models"
)

var _ nsdb.SelectQueryFilter = (*LiveMatchStatsFilters)(nil)

type LiveMatchStatsFilters struct {
	MatchIDs nscol.MatchIDs
	Latest   int
}

func (f LiveMatchStatsFilters) Empty() bool {
	return len(f.MatchIDs) == 0
}

func (f LiveMatchStatsFilters) Validate() error {
	if f.Empty() {
		return xerrors.Errorf("invalid live match stats filters: %w", ErrEmptyFilters)
	}

	return nil
}

func (f LiveMatchStatsFilters) Filter(q *nsdb.SelectQuery) *nsdb.SelectQuery {
	t := nsm.LiveMatchStatsTable
	cMatchID := t.Col("match_id")

	if f.Latest > 0 {
		/*
			SELECT * from live_match_stats
			INNER JOIN (
				SELECT id, row_number() OVER(PARTITION BY match_id ORDER BY created_at DESC) AS row_id
				FROM live_match_stats
				WHERE match_id IN (?)
			) AS latest ON (live_match_stats.id = latest.id and latest.row_id = 1)
		*/

		pk := t.PK()
		tLatest := goqu.T("latest")
		cLatestID := tLatest.Col("id")
		cLatestRowID := tLatest.Col("row_id")

		exprRowID := goqu.ROW_NUMBER().Over(
			goqu.W().
				PartitionBy(cMatchID).
				OrderBy(t.CreatedAt().Desc()),
		)

		// FIXME: qLatest should be created from root, not from `q`
		qLatest := q.
			Select(pk, exprRowID.As(cLatestRowID.GetCol())).
			From(t).
			Where(cMatchID.In(f.MatchIDs))

		qLatestExpr := qLatest.As(tLatest.GetTable()).Expr()

		joinCond := goqu.On(pk.Eq(cLatestID), cLatestRowID.Lte(f.Latest))

		q = q.
			InnerJoin(qLatestExpr, joinCond).
			Prepared(true).
			Trace()
	} else {
		l := len(f.MatchIDs)

		if l == 1 {
			q = q.Eq(cMatchID, f.MatchIDs[0])
		} else {
			q = q.In(cMatchID, f.MatchIDs)
		}

		q = q.Prepared(true)
	}

	return q
}
