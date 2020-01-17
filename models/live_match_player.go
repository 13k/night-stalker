package models

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/paralin/go-dota2/protocol"
)

var LiveMatchPlayerModel = (*LiveMatchPlayer)(nil)

type LiveMatchPlayerID uint64

// LiveMatchPlayer ...
type LiveMatchPlayer struct {
	ID          LiveMatchPlayerID `gorm:"column:id;primary_key"`
	LiveMatchID LiveMatchID       `gorm:"column:live_match_id;unique_index:uix_live_match_players_live_match_id_account_id;not null"` //nolint: lll
	AccountID   nspb.AccountID    `gorm:"column:account_id;unique_index:uix_live_match_players_live_match_id_account_id;not null"`    //nolint: lll
	HeroID      HeroID            `gorm:"column:hero_id"`
	Timestamps

	LiveMatch *LiveMatch
	Hero      *Hero
}

func (*LiveMatchPlayer) TableName() string {
	return "live_match_players"
}

func LiveMatchPlayerDotaProto(pb *protocol.CSourceTVGameSmall_Player) *LiveMatchPlayer {
	return &LiveMatchPlayer{
		AccountID: pb.GetAccountId(),
		HeroID:    HeroID(pb.GetHeroId()),
	}
}
