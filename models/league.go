package models

import (
	"database/sql"

	d2pb "github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
)

var LeagueTable = NewTable("leagues")

type League struct {
	ID `db:"id"`

	Name           string            `db:"name"`
	Tier           nspb.LeagueTier   `db:"tier"`
	Region         nspb.LeagueRegion `db:"region"`
	Status         nspb.LeagueStatus `db:"status"`
	TotalPrizePool uint32            `db:"total_prize_pool"`
	LastActivityAt sql.NullTime      `db:"last_activity_at"`
	StartAt        sql.NullTime      `db:"start_at"`
	FinishAt       sql.NullTime      `db:"finish_at"`

	Timestamps
	SoftDelete
}

func NewLeagueProto(pb *d2pb.CMsgDOTALeagueInfo) *League {
	return &League{
		ID:             ID(pb.GetLeagueId()),
		Name:           pb.GetName(),
		Tier:           nspb.LeagueTier(pb.GetTier()),
		Region:         nspb.LeagueRegion(pb.GetRegion()),
		Status:         nspb.LeagueStatus(pb.GetStatus()),
		TotalPrizePool: pb.GetTotalPrizePool(),
		LastActivityAt: nssql.NullTimeUnix(int64(pb.GetMostRecentActivity())),
		StartAt:        nssql.NullTimeUnix(int64(pb.GetStartTimestamp())),
		FinishAt:       nssql.NullTimeUnix(int64(pb.GetEndTimestamp())),
	}
}
