package dataaccess

import (
	nsdb "github.com/13k/night-stalker/internal/db"
	nsm "github.com/13k/night-stalker/models"
)

var _ nsdb.SelectQueryFilter = (*ProPlayerFilters)(nil)

type ProPlayerFilters struct {
	PlayerMetaFilters
}

func (f ProPlayerFilters) Filter(q *nsdb.SelectQuery) *nsdb.SelectQuery {
	return f.PlayerMetaFilters.Filter(q, nsm.ProPlayerModel)
}
