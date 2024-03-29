package models

import (
	"database/sql"

	d2pb "github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
)

var LiveMatchModel Model = (*LiveMatch)(nil)

type LiveMatchID uint64

// LiveMatch ...
type LiveMatch struct {
	ID                         LiveMatchID        `gorm:"column:id;primary_key"`
	MatchID                    nspb.MatchID       `gorm:"column:match_id;unique_index;not null"`
	ServerID                   nspb.SteamID       `gorm:"column:server_id;not null"`
	LobbyID                    nspb.LobbyID       `gorm:"column:lobby_id;not null"`
	LobbyType                  nspb.LobbyType     `gorm:"column:lobby_type"`
	LeagueID                   nspb.LeagueID      `gorm:"column:league_id"`
	SeriesID                   nspb.SeriesID      `gorm:"column:series_id"`
	GameMode                   nspb.GameMode      `gorm:"column:game_mode"`
	AverageMMR                 uint32             `gorm:"column:average_mmr"`
	RadiantLead                int32              `gorm:"column:radiant_lead"`
	RadiantTeamID              nspb.TeamID        `gorm:"column:radiant_team_id"`
	RadiantTeamName            string             `gorm:"column:radiant_team_name;size:255"`
	RadiantTeamLogo            nspb.SteamID       `gorm:"column:radiant_team_logo"`
	RadiantScore               uint32             `gorm:"column:radiant_score"`
	DireTeamID                 nspb.TeamID        `gorm:"column:dire_team_id"`
	DireTeamName               string             `gorm:"column:dire_team_name;size:255"`
	DireTeamLogo               nspb.SteamID       `gorm:"column:dire_team_logo"`
	DireScore                  uint32             `gorm:"column:dire_score"`
	Delay                      uint32             `gorm:"column:delay"`
	ActivateTime               sql.NullTime       `gorm:"column:activate_time"`
	DeactivateTime             sql.NullTime       `gorm:"column:deactivate_time"`
	LastUpdateTime             sql.NullTime       `gorm:"column:last_update_time"`
	GameTime                   int32              `gorm:"column:game_time"`
	Spectators                 uint32             `gorm:"column:spectators"`
	SortScore                  float64            `gorm:"column:sort_score"`
	BuildingState              nspb.BuildingState `gorm:"column:building_state"`
	WeekendTourneyTournamentID uint32             `gorm:"column:weekend_tourney_tournament_id"`
	WeekendTourneyDivision     uint32             `gorm:"column:weekend_tourney_division"`
	WeekendTourneySkillLevel   uint32             `gorm:"column:weekend_tourney_skill_level"`
	WeekendTourneyBracketRound uint32             `gorm:"column:weekend_tourney_bracket_round"`
	Timestamps

	Players []*LiveMatchPlayer
	Stats   []*LiveMatchStats
	Match   *Match
}

func (*LiveMatch) TableName() string {
	return "live_matches"
}

func (m *LiveMatch) Equal(other *LiveMatch) bool {
	return m.MatchID == other.MatchID &&
		m.ServerID == other.ServerID &&
		m.LobbyID == other.LobbyID &&
		m.LobbyType == other.LobbyType &&
		m.LeagueID == other.LeagueID &&
		m.SeriesID == other.SeriesID &&
		m.GameMode == other.GameMode &&
		m.AverageMMR == other.AverageMMR &&
		m.RadiantLead == other.RadiantLead &&
		m.RadiantTeamID == other.RadiantTeamID &&
		m.RadiantTeamName == other.RadiantTeamName &&
		m.RadiantTeamLogo == other.RadiantTeamLogo &&
		m.RadiantScore == other.RadiantScore &&
		m.DireTeamID == other.DireTeamID &&
		m.DireTeamName == other.DireTeamName &&
		m.DireTeamLogo == other.DireTeamLogo &&
		m.DireScore == other.DireScore &&
		m.Delay == other.Delay &&
		nssql.NullTimeEqual(m.ActivateTime, other.ActivateTime) &&
		nssql.NullTimeEqual(m.DeactivateTime, other.DeactivateTime) &&
		nssql.NullTimeEqual(m.LastUpdateTime, other.LastUpdateTime) &&
		m.GameTime == other.GameTime &&
		m.Spectators == other.Spectators &&
		m.SortScore == other.SortScore &&
		m.BuildingState == other.BuildingState &&
		m.WeekendTourneyTournamentID == other.WeekendTourneyTournamentID &&
		m.WeekendTourneyDivision == other.WeekendTourneyDivision &&
		m.WeekendTourneySkillLevel == other.WeekendTourneySkillLevel &&
		m.WeekendTourneyBracketRound == other.WeekendTourneyBracketRound
}

func LiveMatchDotaProto(pb *d2pb.CSourceTVGameSmall) *LiveMatch {
	return &LiveMatch{
		MatchID:                    nspb.MatchID(pb.GetMatchId()),
		ServerID:                   nspb.SteamID(pb.GetServerSteamId()),
		LeagueID:                   nspb.LeagueID(pb.GetLeagueId()),
		SeriesID:                   nspb.SeriesID(pb.GetSeriesId()),
		LobbyID:                    nspb.LobbyID(pb.GetLobbyId()),
		LobbyType:                  nspb.LobbyType(pb.GetLobbyType()),
		GameMode:                   nspb.GameMode(pb.GetGameMode()),
		AverageMMR:                 pb.GetAverageMmr(),
		RadiantLead:                pb.GetRadiantLead(),
		RadiantTeamID:              nspb.TeamID(pb.GetTeamIdRadiant()),
		RadiantTeamName:            pb.GetTeamNameRadiant(),
		RadiantTeamLogo:            nspb.SteamID(TruncateUint(pb.GetTeamLogoRadiant())),
		RadiantScore:               pb.GetRadiantScore(),
		DireTeamID:                 nspb.TeamID(pb.GetTeamIdDire()),
		DireTeamName:               pb.GetTeamNameDire(),
		DireTeamLogo:               nspb.SteamID(TruncateUint(pb.GetTeamLogoDire())),
		DireScore:                  pb.GetDireScore(),
		ActivateTime:               nssql.NullTimeUnix(int64(pb.GetActivateTime())),
		DeactivateTime:             nssql.NullTimeUnix(int64(pb.GetDeactivateTime())),
		LastUpdateTime:             nssql.NullTimeUnix(int64(pb.GetLastUpdateTime())),
		Delay:                      pb.GetDelay(),
		GameTime:                   pb.GetGameTime(),
		Spectators:                 pb.GetSpectators(),
		SortScore:                  float64(pb.GetSortScore()),
		BuildingState:              nspb.BuildingState(pb.GetBuildingState()),
		WeekendTourneyTournamentID: pb.GetWeekendTourneyTournamentId(),
		WeekendTourneyDivision:     pb.GetWeekendTourneyDivision(),
		WeekendTourneySkillLevel:   pb.GetWeekendTourneySkillLevel(),
		WeekendTourneyBracketRound: pb.GetWeekendTourneyBracketRound(),
	}
}
