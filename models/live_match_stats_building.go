package models

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/paralin/go-dota2/protocol"
)

var LiveMatchStatsBuildingModel = (*LiveMatchStatsBuilding)(nil)

type LiveMatchStatsBuildingID uint64

// LiveMatchStatsBuilding ...
type LiveMatchStatsBuilding struct {
	ID               LiveMatchStatsBuildingID `gorm:"column:id;primary_key"`
	LiveMatchStatsID LiveMatchStatsID         `gorm:"column:live_match_stats_id"`
	GameTeam         nspb.GameTeam            `gorm:"column:game_team"`
	Heading          float32                  `gorm:"column:heading"`
	Type             nspb.BuildingType        `gorm:"column:type"`
	Lane             nspb.LaneType            `gorm:"column:lane"`
	Tier             uint32                   `gorm:"column:tier"`
	PosX             float32                  `gorm:"column:pos_x"`
	PosY             float32                  `gorm:"column:pos_y"`
	Destroyed        bool                     `gorm:"column:destroyed"`
	Timestamps

	LiveMatchStats *LiveMatchStats
}

func (*LiveMatchStatsBuilding) TableName() string {
	return "live_match_stats_buildings"
}

func LiveMatchStatsBuildingDotaProto(
	pb *protocol.CMsgDOTARealtimeGameStatsTerse_BuildingDetails,
) *LiveMatchStatsBuilding {
	return &LiveMatchStatsBuilding{
		GameTeam:  nspb.GameTeam(pb.GetTeam()),
		Heading:   pb.GetHeading(),
		Type:      nspb.BuildingType(pb.GetType()),
		Lane:      nspb.LaneType(pb.GetLane()),
		Tier:      pb.GetTier(),
		PosX:      pb.GetX(),
		PosY:      pb.GetY(),
		Destroyed: pb.GetDestroyed(),
	}
}
