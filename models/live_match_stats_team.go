package models

import (
	"github.com/faceit/go-steam/steamid"
	"github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

var LiveMatchStatsTeamModel Model = (*LiveMatchStatsTeam)(nil)

type LiveMatchStatsTeamID uint64

// LiveMatchStatsTeam ...
type LiveMatchStatsTeam struct {
	ID               LiveMatchStatsTeamID `gorm:"column:id;primary_key"`
	LiveMatchStatsID LiveMatchStatsID     `gorm:"column:live_match_stats_id"`
	TeamID           TeamID               `gorm:"column:team_id"`
	GameTeam         nspb.GameTeam        `gorm:"column:game_team"`
	Name             string               `gorm:"column:name;size:255"`
	Tag              string               `gorm:"column:tag;size:255"`
	LogoID           steamid.SteamId      `gorm:"column:logo_id"`
	LogoURL          string               `gorm:"column:logo_url"`
	Score            uint32               `gorm:"column:score"`
	NetWorth         uint32               `gorm:"column:net_worth"`
	Timestamps

	LiveMatchStats *LiveMatchStats
	Team           *Team
}

func (*LiveMatchStatsTeam) TableName() string {
	return "live_match_stats_teams"
}

func LiveMatchStatsTeamDotaProto(pb *protocol.CMsgDOTARealtimeGameStatsTerse_TeamDetails) *LiveMatchStatsTeam {
	return &LiveMatchStatsTeam{
		TeamID:   TeamID(pb.GetTeamId()),
		GameTeam: nspb.GameTeam(pb.GetTeamNumber()),
		Name:     pb.GetTeamName(),
		Tag:      pb.GetTeamTag(),
		LogoID:   steamid.SteamId(pb.GetTeamLogo()),
		LogoURL:  pb.GetTeamLogoUrl(),
		Score:    pb.GetScore(),
		NetWorth: pb.GetNetWorth(),
	}
}
