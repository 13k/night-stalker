package dataaccess

import (
	"time"

	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsdb "github.com/13k/night-stalker/internal/db"
	nsm "github.com/13k/night-stalker/models"
)

var _ nsdb.SelectQueryFilter = (*MatchFilters)(nil)

type MatchFilters struct {
	MatchIDs    nscol.MatchIDs
	MinDuration time.Duration
	MaxDuration time.Duration
	Since       time.Time
	Players     PlayerStatsFilters
}

func (f MatchFilters) Empty() bool {
	return len(f.MatchIDs) == 0 &&
		f.MinDuration == 0 &&
		f.MaxDuration == 0 &&
		f.Since.IsZero() &&
		f.Players.Empty()
}

func (f MatchFilters) Validate() error {
	if f.Empty() {
		return xerrors.Errorf("invalid match filters: %w", ErrEmptyFilters)
	}

	return nil
}

func (f MatchFilters) Filter(q *nsdb.SelectQuery) *nsdb.SelectQuery {
	t := nsm.MatchTable

	if l := len(f.MatchIDs); l > 0 {
		pk := t.PK()

		if l == 1 {
			q = q.Eq(pk, f.MatchIDs[0])
		} else {
			q = q.In(pk, f.MatchIDs)
		}

		q = q.Prepared(true)
	}

	colDuration := t.Col("duration")

	if f.MinDuration != 0 {
		q = q.
			GtEq(colDuration, int64(f.MinDuration/time.Second)).
			Prepared(true)
	}

	if f.MaxDuration != 0 {
		q = q.
			LtEq(colDuration, int64(f.MaxDuration/time.Second)).
			Prepared(true)
	}

	if !f.Since.IsZero() {
		q = q.
			GtEq(t.CreatedAt(), f.Since).
			Prepared(true)
	}

	return q
}
