package dataaccess

import (
	"golang.org/x/xerrors"

	nscol "github.com/13k/night-stalker/internal/collections"
	nsdb "github.com/13k/night-stalker/internal/db"
	nsm "github.com/13k/night-stalker/models"
)

// PlayerMetaFilters holds filters for `FollowedPlayer`, `Player` and `ProPlayer`.
//
// It cannot be used directly as a `QueryFilterer` but can be used by one.
type PlayerMetaFilters struct {
	AccountIDs nscol.AccountIDs
	Query      string
}

func (f PlayerMetaFilters) Empty() bool {
	return len(f.AccountIDs) == 0 && f.Query == ""
}

func (f PlayerMetaFilters) Validate() error {
	if f.Empty() {
		return xerrors.Errorf("invalid player filters: %w", ErrEmptyFilters)
	}

	return nil
}

func (f PlayerMetaFilters) Filter(q *nsdb.SelectQuery, m nsm.Model) *nsdb.SelectQuery {
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

	if f.Query != "" {
		likePattern := "%" + f.Query + "%"

		switch t {
		case nsm.FollowedPlayerTable:
			q = q.ILike(t.Col("label"), likePattern)
		case nsm.PlayerTable:
			q = q.ILike(t.Col("name"), likePattern)
		case nsm.ProPlayerTable:
			// not supported
		}

		q = q.Prepared(true)
	}

	return q
}
