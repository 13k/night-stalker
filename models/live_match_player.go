package models

import (
	"github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

var LiveMatchPlayerModel Model = (*LiveMatchPlayer)(nil)

type LiveMatchPlayerID uint64

// LiveMatchPlayer ...
type LiveMatchPlayer struct {
	ID          LiveMatchPlayerID `gorm:"column:id;primary_key"`
	LiveMatchID LiveMatchID       `gorm:"column:live_match_id;unique_index:uix_live_match_players_live_match_id_account_id;not null"` //nolint: lll
	MatchID     nspb.MatchID      `gorm:"column:match_id;unique_index:uix_live_match_players_match_id_account_id;not null"`           //nolint: lll
	AccountID   nspb.AccountID    `gorm:"column:account_id;unique_index:uix_live_match_players_live_match_id_account_id;not null"`    //nolint: lll
	HeroID      nspb.HeroID       `gorm:"column:hero_id"`
	Timestamps

	LiveMatch *LiveMatch
	Match     *Match
	Hero      *Hero
}

func (*LiveMatchPlayer) TableName() string {
	return "live_match_players"
}

func NewLiveMatchPlayer(liveMatch *LiveMatch, pb *protocol.CSourceTVGameSmall_Player) *LiveMatchPlayer {
	p := LiveMatchPlayerDotaProto(pb)
	p.LiveMatchID = liveMatch.ID
	p.MatchID = liveMatch.MatchID
	return p
}

func LiveMatchPlayerDotaProto(pb *protocol.CSourceTVGameSmall_Player) *LiveMatchPlayer {
	return &LiveMatchPlayer{
		AccountID: nspb.AccountID(pb.GetAccountId()),
		HeroID:    nspb.HeroID(pb.GetHeroId()),
	}
}
