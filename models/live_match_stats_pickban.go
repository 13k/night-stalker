package models

import (
	d2pb "github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

var LiveMatchStatsPickBanTable = NewTable("live_match_stats_picksbans")

type LiveMatchStatsPickBan struct {
	ID `db:"id" goqu:"defaultifempty"`

	GameTeam nspb.GameTeam `db:"game_team"`
	IsBan    bool          `db:"is_ban"`

	LiveMatchStatsID ID `db:"live_match_stats_id"`
	HeroID           ID `db:"hero_id"`

	Timestamps
	SoftDelete

	LiveMatchStats *LiveMatchStats `db:"-" model:"belongs_to"`
	Hero           *Hero           `db:"-" model:"belongs_to"`
}

func LiveMatchStatsPickBanAssocProto(
	liveMatchStats *LiveMatchStats,
	isBan bool,
	pb *d2pb.CMsgDOTARealtimeGameStatsTerse_PickBanDetails,
) *LiveMatchStatsPickBan {
	m := NewLiveMatchStatsPickBanProto(isBan, pb)
	m.LiveMatchStats = liveMatchStats
	m.LiveMatchStatsID = liveMatchStats.ID
	return m
}

func NewLiveMatchStatsPickBanProto(
	isBan bool,
	pb *d2pb.CMsgDOTARealtimeGameStatsTerse_PickBanDetails,
) *LiveMatchStatsPickBan {
	return &LiveMatchStatsPickBan{
		HeroID:   ID(pb.GetHero()),
		GameTeam: nspb.GameTeam(pb.GetTeam()),
		IsBan:    isBan,
	}
}
