package dataaccess

import (
	nsdb "github.com/13k/night-stalker/internal/db"
	nsm "github.com/13k/night-stalker/models"
)

var _ nsdb.SelectQueryFilter = (*PlayerFilters)(nil)

type PlayerFilters struct {
	PlayerMetaFilters
}

func (f PlayerFilters) Filter(q *nsdb.SelectQuery) *nsdb.SelectQuery {
	return f.PlayerMetaFilters.Filter(q, nsm.PlayerModel)
}
