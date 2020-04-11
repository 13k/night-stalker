package models

import (
	d2pb "github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
)

var LiveMatchStatsPlayerTable = NewTable("live_match_stats_players")

type LiveMatchStatsPlayer struct {
	ID `db:"id" goqu:"defaultifempty"`

	AccountID  nspb.AccountID      `db:"account_id"`
	PlayerSlot nspb.GamePlayerSlot `db:"player_slot"`
	Name       string              `db:"name"`
	GameTeam   nspb.GameTeam       `db:"game_team"`
	Level      uint32              `db:"level"`
	Kills      uint32              `db:"kills"`
	Deaths     uint32              `db:"deaths"`
	Assists    uint32              `db:"assists"`
	Denies     uint32              `db:"denies"`
	LastHits   uint32              `db:"last_hits"`
	Gold       uint32              `db:"gold"`
	NetWorth   uint32              `db:"net_worth"`
	PosX       float32             `db:"pos_x"`
	PosY       float32             `db:"pos_y"`
	Abilities  nssql.AbilityIDs    `db:"abilities"`
	Items      nssql.ItemIDs       `db:"items"`

	LiveMatchStatsID ID `db:"live_match_stats_id"`
	MatchID          ID `db:"match_id"`
	HeroID           ID `db:"hero_id"`

	Timestamps
	SoftDelete

	LiveMatchStats *LiveMatchStats `db:"-" model:"belongs_to"`
	Match          *Match          `db:"-" model:"belongs_to"`
	Hero           *Hero           `db:"-" model:"belongs_to"`
}

func NewLiveMatchStatsPlayerAssocProto(
	liveMatchStats *LiveMatchStats,
	pb *d2pb.CMsgDOTARealtimeGameStatsTerse_PlayerDetails,
) *LiveMatchStatsPlayer {
	m := NewLiveMatchStatsPlayerProto(pb)
	m.LiveMatchStats = liveMatchStats
	m.LiveMatchStatsID = liveMatchStats.ID
	m.MatchID = liveMatchStats.MatchID
	return m
}

func NewLiveMatchStatsPlayerProto(pb *d2pb.CMsgDOTARealtimeGameStatsTerse_PlayerDetails) *LiveMatchStatsPlayer {
	return &LiveMatchStatsPlayer{
		AccountID:  nspb.AccountID(pb.GetAccountid()),
		PlayerSlot: nspb.GamePlayerIndex(pb.GetPlayerid()).GamePlayerSlot(),
		Name:       pb.GetName(),
		GameTeam:   nspb.GameTeam(pb.GetTeam()),
		HeroID:     ID(pb.GetHeroid()),
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
		Abilities:  nssql.NewAbilityIDs(pb.GetAbilities()),
		Items:      nssql.NewItemIDs(pb.GetItems()),
	}
}
