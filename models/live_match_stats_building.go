package models

import (
	d2pb "github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

var LiveMatchStatsBuildingTable = NewTable("live_match_stats_buildings")

type LiveMatchStatsBuilding struct {
	ID `db:"id" goqu:"defaultifempty"`

	GameTeam  nspb.GameTeam     `db:"game_team"`
	Heading   float32           `db:"heading"`
	Type      nspb.BuildingType `db:"type"`
	Lane      nspb.LaneType     `db:"lane"`
	Tier      uint32            `db:"tier"`
	PosX      float32           `db:"pos_x"`
	PosY      float32           `db:"pos_y"`
	Destroyed bool              `db:"destroyed"`

	LiveMatchStatsID ID `db:"live_match_stats_id"`

	Timestamps
	SoftDelete

	LiveMatchStats *LiveMatchStats `db:"-" model:"belongs_to"`
}

func NewLiveMatchStatsBuildingAssocProto(
	liveMatchStats *LiveMatchStats,
	pb *d2pb.CMsgDOTARealtimeGameStatsTerse_BuildingDetails,
) *LiveMatchStatsBuilding {
	m := NewLiveMatchStatsBuildingProto(pb)
	m.LiveMatchStats = liveMatchStats
	m.LiveMatchStatsID = liveMatchStats.ID
	return m
}

func NewLiveMatchStatsBuildingProto(
	pb *d2pb.CMsgDOTARealtimeGameStatsTerse_BuildingDetails,
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
