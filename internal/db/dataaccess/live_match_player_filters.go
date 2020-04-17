package dataaccess

import (
	nsdb "github.com/13k/night-stalker/internal/db"
	nsm "github.com/13k/night-stalker/models"
)

var _ nsdb.SelectQueryFilter = (*LiveMatchPlayerFilters)(nil)

type LiveMatchPlayerFilters struct {
	PlayerStatsFilters
}

func (f LiveMatchPlayerFilters) Filter(q *nsdb.SelectQuery) *nsdb.SelectQuery {
	return f.PlayerStatsFilters.Filter(q, nsm.LiveMatchPlayerModel)
}
