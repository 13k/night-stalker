package models

import (
	d2pb "github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
)

var MatchPlayerTable = NewTable("match_players")

type MatchPlayer struct {
	ID `db:"id" goqu:"defaultifempty"`

	AccountID  nspb.AccountID      `db:"account_id"`
	PlayerSlot nspb.GamePlayerSlot `db:"player_slot"`
	ProName    string              `db:"pro_name"`
	Kills      uint32              `db:"kills"`
	Deaths     uint32              `db:"deaths"`
	Assists    uint32              `db:"assists"`
	Items      nssql.ItemIDs       `db:"items"`

	MatchID ID `db:"match_id"`
	HeroID  ID `db:"hero_id"`

	Timestamps
	SoftDelete

	Match *Match `db:"-" model:"belongs_to"`
	Hero  *Hero  `db:"-" model:"belongs_to"`
}

func NewMatchPlayerAssocProto(
	match *Match,
	pb *d2pb.CMsgDOTAMatchMinimal_Player,
) *MatchPlayer {
	m := NewMatchPlayerProto(pb)
	m.Match = match
	m.MatchID = match.ID
	return m
}

func NewMatchPlayerProto(pb *d2pb.CMsgDOTAMatchMinimal_Player) *MatchPlayer {
	return &MatchPlayer{
		AccountID:  nspb.AccountID(pb.GetAccountId()),
		HeroID:     ID(pb.GetHeroId()),
		PlayerSlot: nspb.GamePlayerSlot(pb.GetPlayerSlot()),
		ProName:    pb.GetProName(),
		Kills:      pb.GetKills(),
		Deaths:     pb.GetDeaths(),
		Assists:    pb.GetAssists(),
		Items:      nssql.NewItemIDs(pb.GetItems()),
	}
}
