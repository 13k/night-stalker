package models

import (
	nsproto "github.com/13k/night-stalker/internal/protocol"
	"github.com/lib/pq"
	"github.com/paralin/go-dota2/protocol"
)

var LiveMatchStatsPlayerModel = (*LiveMatchStatsPlayer)(nil)

type LiveMatchStatsPlayerID uint64

// LiveMatchStatsPlayer ...
type LiveMatchStatsPlayer struct {
	ID               LiveMatchStatsPlayerID `gorm:"column:id;primary_key"`
	LiveMatchStatsID LiveMatchStatsID       `gorm:"column:live_match_stats_id"`
	AccountID        nsproto.AccountID      `gorm:"column:account_id"`
	PlayerSlot       nsproto.GamePlayerSlot `gorm:"column:player_slot"`
	Name             string                 `gorm:"column:name;size:255"`
	GameTeam         nsproto.GameTeam       `gorm:"column:game_team"`
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

func LiveMatchStatsPlayerDotaProto(pb *protocol.CMsgDOTARealtimeGameStatsTerse_PlayerDetails) *LiveMatchStatsPlayer {
	return &LiveMatchStatsPlayer{
		AccountID:  pb.GetAccountid(),
		PlayerSlot: nsproto.GamePlayerSlot(pb.GetPlayerid()),
		Name:       pb.GetName(),
		GameTeam:   nsproto.GameTeam(pb.GetTeam()),
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
