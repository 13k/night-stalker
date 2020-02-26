package models

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/lib/pq"
	"github.com/paralin/go-dota2/protocol"
)

var LiveMatchStatsPlayerModel = (*LiveMatchStatsPlayer)(nil)

type LiveMatchStatsPlayerID uint64

// LiveMatchStatsPlayer ...
type LiveMatchStatsPlayer struct {
	ID               LiveMatchStatsPlayerID `gorm:"column:id;primary_key"`
	LiveMatchStatsID LiveMatchStatsID       `gorm:"column:live_match_stats_id;unique_index:uix_live_match_stats_players_live_match_stats_id_account_id;not null"` //nolint: lll
	MatchID          nspb.MatchID           `gorm:"column:match_id;index;not null"`                                                                               //nolint: lll
	AccountID        nspb.AccountID         `gorm:"column:account_id;unique_index:uix_live_match_stats_players_live_match_stats_id_account_id;not null"`          //nolint: lll
	PlayerSlot       nspb.GamePlayerSlot    `gorm:"column:player_slot"`
	Name             string                 `gorm:"column:name;size:255"`
	GameTeam         nspb.GameTeam          `gorm:"column:game_team"`
	HeroID           HeroID                 `gorm:"column:hero_id"`
	Level            uint32                 `gorm:"column:level"`
	Kills            uint32                 `gorm:"column:kills"`
	Deaths           uint32                 `gorm:"column:deaths"`
	Assists          uint32                 `gorm:"column:assists"`
	Denies           uint32                 `gorm:"column:denies"`
	LastHits         uint32                 `gorm:"column:last_hits"`
	Gold             uint32                 `gorm:"column:gold"`
	NetWorth         uint32                 `gorm:"column:net_worth"`
	PosX             float32                `gorm:"column:pos_x"`
	PosY             float32                `gorm:"column:pos_y"`
	Abilities        pq.Int64Array          `gorm:"column:abilities"`
	Items            pq.Int64Array          `gorm:"column:items"`
	Timestamps

	LiveMatchStats *LiveMatchStats
	Hero           *Hero
}

func (*LiveMatchStatsPlayer) TableName() string {
	return "live_match_stats_players"
}

func NewLiveMatchStatsPlayer(
	liveMatchStats *LiveMatchStats,
	pb *protocol.CMsgDOTARealtimeGameStatsTerse_PlayerDetails,
) *LiveMatchStatsPlayer {
	p := LiveMatchStatsPlayerDotaProto(pb)
	p.LiveMatchStatsID = liveMatchStats.ID
	p.MatchID = liveMatchStats.MatchID
	return p
}

func LiveMatchStatsPlayerDotaProto(pb *protocol.CMsgDOTARealtimeGameStatsTerse_PlayerDetails) *LiveMatchStatsPlayer {
	return &LiveMatchStatsPlayer{
		AccountID:  pb.GetAccountid(),
		PlayerSlot: nspb.GamePlayerIndex(pb.GetPlayerid()).GamePlayerSlot(),
		Name:       pb.GetName(),
		GameTeam:   nspb.GameTeam(pb.GetTeam()),
		HeroID:     HeroID(pb.GetHeroid()),
		Level:      pb.GetLevel(),
		Kills:      pb.GetKillCount(),
		Deaths:     pb.GetDeathCount(),
		Assists:    pb.GetAssistsCount(),
		Denies:     pb.GetDeniesCount(),
		LastHits:   pb.GetLhCount(),
		Gold:       pb.GetGold(),
		PosX:       pb.GetX(),
		PosY:       pb.GetY(),
		NetWorth:   pb.GetNetWorth(),
		Abilities:  Uint32Array(pb.GetAbilities()),
		Items:      Uint32Array(pb.GetItems()),
	}
}
