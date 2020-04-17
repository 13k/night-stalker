package dataaccess

import (
	nsdb "github.com/13k/night-stalker/internal/db"
	nsm "github.com/13k/night-stalker/models"
)

var _ nsdb.SelectQueryFilter = (*MatchPlayerFilters)(nil)

type MatchPlayerFilters struct {
	PlayerStatsFilters
}

func (f MatchPlayerFilters) Filter(q *nsdb.SelectQuery) *nsdb.SelectQuery {
	return f.PlayerStatsFilters.Filter(q, nsm.MatchPlayerModel)
}
