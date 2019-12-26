package models

import (
	nsproto "github.com/13k/night-stalker/internal/protocol"
	"github.com/paralin/go-dota2/protocol"
)

var LiveMatchStatsBuildingModel = (*LiveMatchStatsBuilding)(nil)

type LiveMatchStatsBuildingID uint64

// LiveMatchStatsBuilding ...
type LiveMatchStatsBuilding struct {
	ID               LiveMatchStatsBuildingID `gorm:"column:id;primary_key"`
	LiveMatchStatsID LiveMatchStatsID         `gorm:"column:live_match_stats_id"`
	GameTeam         nsproto.GameTeam         `gorm:"column:game_team"`
	Heading          float32                  `gorm:"column:heading"`
	Type             nsproto.BuildingType     `gorm:"column:type"`
	Lane             nsproto.LaneType         `gorm:"column:lane"`
	Tier             uint32                   `gorm:"column:tier"`
	PosX             float32                  `gorm:"column:pos_x"`
	PosY             float32                  `gorm:"column:pos_y"`
	Destroyed        bool                     `gorm:"column:destroyed"`
	Timestamps

	LiveMatchStats *LiveMatchStats
}

func LiveMatchStatsBuildingDotaProto(
	pb *protocol.CMsgDOTARealtimeGameStatsTerse_BuildingDetails,
) *LiveMatchStatsBuilding {
	return &LiveMatchStatsBuilding{
		GameTeam:  nsproto.GameTeam(pb.GetTeam()),
		Heading:   pb.GetHeading(),
		Type:      nsproto.BuildingType(pb.GetType()),
		Lane:      nsproto.LaneType(pb.GetLane()),
		Tier:      pb.GetTier(),
		PosX:      pb.GetX(),
		PosY:      pb.GetY(),
		Destroyed: pb.GetDestroyed(),
	}
}
