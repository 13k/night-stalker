package models

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/lib/pq"
	"github.com/paralin/go-dota2/protocol"
)

var MatchPlayerModel = (*MatchPlayer)(nil)

type MatchPlayerID uint64

// MatchPlayer ...
type MatchPlayer struct {
	ID         MatchPlayerID       `gorm:"column:id;primary_key"`
	MatchID    nspb.MatchID        `gorm:"column:match_id;unique_index:uix_match_players_match_id_account_id;not null"`   //nolint: lll
	AccountID  nspb.AccountID      `gorm:"column:account_id;unique_index:uix_match_players_match_id_account_id;not null"` //nolint: lll
	HeroID     HeroID              `gorm:"column:hero_id"`
	PlayerSlot nspb.GamePlayerSlot `gorm:"column:player_slot"`
	ProName    string              `gorm:"column:pro_name"`
	Kills      uint32              `gorm:"column:kills"`
	Deaths     uint32              `gorm:"column:deaths"`
	Assists    uint32              `gorm:"column:assists"`
	Items      pq.Int64Array       `gorm:"column:items"`
	Timestamps

	Match *Match
	Hero  *Hero
}

func (*MatchPlayer) TableName() string {
	return "match_players"
}

func MatchPlayerDotaProto(pb *protocol.CMsgDOTAMatchMinimal_Player) *MatchPlayer {
	return &MatchPlayer{
		AccountID:  pb.GetAccountId(),
		HeroID:     HeroID(pb.GetHeroId()),
		PlayerSlot: nspb.GamePlayerSlot(pb.GetPlayerSlot()),
		ProName:    pb.GetProName(),
		Kills:      pb.GetKills(),
		Deaths:     pb.GetDeaths(),
		Assists:    pb.GetAssists(),
		Items:      Uint32Array(pb.GetItems()),
	}
}
