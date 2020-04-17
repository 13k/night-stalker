package dataaccess

import (
	nsdb "github.com/13k/night-stalker/internal/db"
	nsm "github.com/13k/night-stalker/models"
)

var _ nsdb.SelectQueryFilter = (*FollowedPlayerFilters)(nil)

type FollowedPlayerFilters struct {
	PlayerMetaFilters
}

func (f FollowedPlayerFilters) Filter(q *nsdb.SelectQuery) *nsdb.SelectQuery {
	return f.PlayerMetaFilters.Filter(q, nsm.FollowedPlayerModel)
}
