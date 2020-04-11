package models

import (
	"database/sql"

	d2pb "github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
)

var MatchTable = NewTable("matches")

type Match struct {
	ID `db:"id"`

	SeriesType                   nspb.SeriesType   `db:"series_type"`
	SeriesGame                   uint32            `db:"series_game"`
	GameMode                     nspb.GameMode     `db:"game_mode"`
	StartTime                    sql.NullTime      `db:"start_time"`
	Duration                     uint32            `db:"duration"`
	Outcome                      nspb.MatchOutcome `db:"outcome"`
	RadiantTeamName              string            `db:"radiant_team_name"`
	RadiantTeamLogo              nspb.SteamID      `db:"radiant_team_logo"`
	RadiantTeamLogoURL           string            `db:"radiant_team_logo_url"`
	RadiantScore                 uint32            `db:"radiant_score"`
	DireTeamName                 string            `db:"dire_team_name"`
	DireTeamLogo                 nspb.SteamID      `db:"dire_team_logo"`
	DireTeamLogoURL              string            `db:"dire_team_logo_url"`
	DireScore                    uint32            `db:"dire_score"`
	WeekendTourneyTournamentID   uint32            `db:"weekend_tourney_tournament_id"`
	WeekendTourneySeasonTrophyID uint32            `db:"weekend_tourney_season_trophy_id"`
	WeekendTourneyDivision       uint32            `db:"weekend_tourney_division"`
	WeekendTourneySkillLevel     uint32            `db:"weekend_tourney_skill_level"`

	LeagueID      ID `db:"league_id"`
	RadiantTeamID ID `db:"radiant_team_id"`
	DireTeamID    ID `db:"dire_team_id"`

	Timestamps
	SoftDelete

	League      *League        `db:"-" model:"belongs_to"`
	RadiantTeam *Team          `db:"-" model:"belongs_to"`
	DireTeam    *Team          `db:"-" model:"belongs_to"`
	LiveMatch   *LiveMatch     `db:"-" model:"has_one"`
	Players     []*MatchPlayer `db:"-" model:"has_many"`
}

func NewMatchProto(pb *d2pb.CMsgDOTAMatchMinimal) *Match {
	return &Match{
		ID:                           ID(pb.GetMatchId()),
		LeagueID:                     ID(pb.GetTourney().GetLeagueId()),
		SeriesType:                   nspb.SeriesType(pb.GetTourney().GetSeriesType()),
		SeriesGame:                   pb.GetTourney().GetSeriesGame(),
		GameMode:                     nspb.GameMode(pb.GetGameMode()),
		StartTime:                    nssql.NullTimeUnix(int64(pb.GetStartTime())),
		Duration:                     pb.GetDuration(),
		Outcome:                      nspb.MatchOutcome(pb.GetMatchOutcome()),
		RadiantTeamID:                ID(pb.GetTourney().GetRadiantTeamId()),
		RadiantTeamName:              pb.GetTourney().GetRadiantTeamName(),
		RadiantTeamLogo:              nspb.SteamID(TruncateUint(pb.GetTourney().GetDireTeamLogo())),
		RadiantTeamLogoURL:           pb.GetTourney().GetRadiantTeamLogoUrl(),
		RadiantScore:                 pb.GetRadiantScore(),
		DireTeamID:                   ID(pb.GetTourney().GetDireTeamId()),
		DireTeamName:                 pb.GetTourney().GetDireTeamName(),
		DireTeamLogo:                 nspb.SteamID(TruncateUint(pb.GetTourney().GetDireTeamLogo())),
		DireTeamLogoURL:              pb.GetTourney().GetDireTeamLogoUrl(),
		DireScore:                    pb.GetDireScore(),
		WeekendTourneyTournamentID:   pb.GetTourney().GetWeekendTourneyTournamentId(),
		WeekendTourneySeasonTrophyID: pb.GetTourney().GetWeekendTourneySeasonTrophyId(),
		WeekendTourneyDivision:       pb.GetTourney().GetWeekendTourneyDivision(),
		WeekendTourneySkillLevel:     pb.GetTourney().GetWeekendTourneySkillLevel(),
	}
}
