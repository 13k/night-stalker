package models

import (
	"time"

	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/faceit/go-steam/steamid"
	"github.com/paralin/go-dota2/protocol"
)

var MatchModel = (*Match)(nil)

// Match ...
type Match struct {
	ID                           nspb.MatchID      `gorm:"column:id;primary_key"`
	LeagueID                     nspb.LeagueID     `gorm:"column:league_id"`
	SeriesType                   uint32            `gorm:"column:series_type"`
	SeriesGame                   uint32            `gorm:"column:series_game"`
	GameMode                     nspb.GameMode     `gorm:"column:game_mode"`
	StartTime                    *time.Time        `gorm:"column:start_time"`
	Duration                     uint32            `gorm:"column:duration"`
	Outcome                      nspb.MatchOutcome `gorm:"column:outcome"`
	RadiantTeamID                TeamID            `gorm:"column:radiant_team_id"`
	RadiantTeamName              string            `gorm:"column:radiant_team_name;size:255"`
	RadiantTeamLogo              steamid.SteamId   `gorm:"column:radiant_team_logo"`
	RadiantTeamLogoURL           string            `gorm:"column:radiant_team_logo_url"`
	RadiantScore                 uint32            `gorm:"column:radiant_score"`
	DireTeamID                   TeamID            `gorm:"column:dire_team_id"`
	DireTeamName                 string            `gorm:"column:dire_team_name;size:255"`
	DireTeamLogo                 steamid.SteamId   `gorm:"column:dire_team_logo"`
	DireTeamLogoURL              string            `gorm:"column:dire_team_logo_url"`
	DireScore                    uint32            `gorm:"column:dire_score"`
	WeekendTourneyTournamentID   uint32            `gorm:"column:weekend_tourney_tournament_id"`
	WeekendTourneySeasonTrophyID uint32            `gorm:"column:weekend_tourney_season_trophy_id"`
	WeekendTourneyDivision       uint32            `gorm:"column:weekend_tourney_division"`
	WeekendTourneySkillLevel     uint32            `gorm:"column:weekend_tourney_skill_level"`
	Timestamps

	Players []*MatchPlayer
}

func (*Match) TableName() string {
	return "matches"
}

func MatchDotaProto(pb *protocol.CMsgDOTAMatchMinimal) *Match {
	return &Match{
		ID:                           pb.GetMatchId(),
		LeagueID:                     nspb.LeagueID(pb.GetTourney().GetLeagueId()),
		SeriesType:                   pb.GetTourney().GetSeriesType(),
		SeriesGame:                   pb.GetTourney().GetSeriesGame(),
		GameMode:                     nspb.GameMode(pb.GetGameMode()),
		StartTime:                    NullUnixTimestamp(int64(pb.GetStartTime())),
		Duration:                     pb.GetDuration(),
		Outcome:                      nspb.MatchOutcome(pb.GetMatchOutcome()),
		RadiantTeamID:                TeamID(pb.GetTourney().GetRadiantTeamId()),
		RadiantTeamName:              pb.GetTourney().GetRadiantTeamName(),
		RadiantTeamLogo:              steamid.SteamId(TruncateUint(pb.GetTourney().GetDireTeamLogo())),
		RadiantTeamLogoURL:           pb.GetTourney().GetRadiantTeamLogoUrl(),
		RadiantScore:                 pb.GetRadiantScore(),
		DireTeamID:                   TeamID(pb.GetTourney().GetDireTeamId()),
		DireTeamName:                 pb.GetTourney().GetDireTeamName(),
		DireTeamLogo:                 steamid.SteamId(TruncateUint(pb.GetTourney().GetDireTeamLogo())),
		DireTeamLogoURL:              pb.GetTourney().GetDireTeamLogoUrl(),
		DireScore:                    pb.GetDireScore(),
		WeekendTourneyTournamentID:   pb.GetTourney().GetWeekendTourneyTournamentId(),
		WeekendTourneySeasonTrophyID: pb.GetTourney().GetWeekendTourneySeasonTrophyId(),
		WeekendTourneyDivision:       pb.GetTourney().GetWeekendTourneyDivision(),
		WeekendTourneySkillLevel:     pb.GetTourney().GetWeekendTourneySkillLevel(),
	}
}
