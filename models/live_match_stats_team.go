package models

import (
	d2pb "github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

var LiveMatchStatsTeamTable = NewTable("live_match_stats_teams")

type LiveMatchStatsTeam struct {
	ID `db:"id" goqu:"defaultifempty"`

	GameTeam nspb.GameTeam `db:"game_team"`
	Name     string        `db:"name"`
	Tag      string        `db:"tag"`
	LogoID   nspb.SteamID  `db:"logo_id"`
	LogoURL  string        `db:"logo_url"`
	Score    uint32        `db:"score"`
	NetWorth uint32        `db:"net_worth"`

	LiveMatchStatsID ID `db:"live_match_stats_id"`
	TeamID           ID `db:"team_id"`

	Timestamps
	SoftDelete

	LiveMatchStats *LiveMatchStats `db:"-" model:"belongs_to,source=Teams"`
	Team           *Team           `db:"-" model:"belongs_to"`
}

func NewLiveMatchStatsTeamAssocProto(
	liveMatchStats *LiveMatchStats,
	pb *d2pb.CMsgDOTARealtimeGameStatsTerse_TeamDetails,
) *LiveMatchStatsTeam {
	m := NewLiveMatchStatsTeamProto(pb)
	m.LiveMatchStats = liveMatchStats
	m.LiveMatchStatsID = liveMatchStats.ID
	return m
}

func NewLiveMatchStatsTeamProto(
	pb *d2pb.CMsgDOTARealtimeGameStatsTerse_TeamDetails,
) *LiveMatchStatsTeam {
	return &LiveMatchStatsTeam{
		TeamID:   ID(pb.GetTeamId()),
		GameTeam: nspb.GameTeam(pb.GetTeamNumber()),
		Name:     pb.GetTeamName(),
		Tag:      pb.GetTeamTag(),
		LogoID:   nspb.SteamID(pb.GetTeamLogo()),
		LogoURL:  pb.GetTeamLogoUrl(),
		Score:    pb.GetScore(),
		NetWorth: pb.GetNetWorth(),
	}
}
