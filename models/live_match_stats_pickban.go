package models

import (
	"github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

var LiveMatchStatsPickBanModel Model = (*LiveMatchStatsPickBan)(nil)

type LiveMatchStatsPickBanID uint64

// LiveMatchStatsPickBan ...
type LiveMatchStatsPickBan struct {
	ID               LiveMatchStatsPickBanID `gorm:"column:id;primary_key"`
	LiveMatchStatsID LiveMatchStatsID        `gorm:"column:live_match_stats_id"`
	HeroID           nspb.HeroID             `gorm:"column:hero_id"`
	GameTeam         nspb.GameTeam           `gorm:"column:game_team"`
	IsBan            bool                    `gorm:"column:is_ban"`
	Timestamps

	LiveMatchStats *LiveMatchStats
	Hero           *Hero
}

func (LiveMatchStatsPickBan) TableName() string {
	return "live_match_stats_picksbans"
}

func LiveMatchStatsPickBanDotaProto(
	isBan bool,
	pb *protocol.CMsgDOTARealtimeGameStatsTerse_PickBanDetails,
) *LiveMatchStatsPickBan {
	return &LiveMatchStatsPickBan{
		HeroID:   nspb.HeroID(pb.GetHero()),
		GameTeam: nspb.GameTeam(pb.GetTeam()),
		IsBan:    isBan,
	}
}
