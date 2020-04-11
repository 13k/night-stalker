package models

import (
	d2pb "github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

var LiveMatchPlayerTable = NewTable("live_match_players")

type LiveMatchPlayer struct {
	ID `db:"id" goqu:"defaultifempty"`

	AccountID nspb.AccountID `db:"account_id"`

	LiveMatchID ID `db:"live_match_id"`
	MatchID     ID `db:"match_id"`
	HeroID      ID `db:"hero_id"`

	Timestamps
	SoftDelete

	LiveMatch *LiveMatch `db:"-" model:"belongs_to"`
	Match     *Match     `db:"-" model:"belongs_to"`
	Hero      *Hero      `db:"-" model:"belongs_to"`
}

func NewLiveMatchPlayerAssocProto(
	liveMatch *LiveMatch,
	pb *d2pb.CSourceTVGameSmall_Player,
) *LiveMatchPlayer {
	m := NewLiveMatchPlayerProto(pb)
	m.LiveMatch = liveMatch
	m.LiveMatchID = liveMatch.ID
	m.MatchID = liveMatch.MatchID
	return m
}

func NewLiveMatchPlayerProto(pb *d2pb.CSourceTVGameSmall_Player) *LiveMatchPlayer {
	return &LiveMatchPlayer{
		AccountID: nspb.AccountID(pb.GetAccountId()),
		HeroID:    ID(pb.GetHeroId()),
	}
}
