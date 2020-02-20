package models

import (
	"time"

	"github.com/faceit/go-steam/steamid"
	"github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protocol"
	nstime "github.com/13k/night-stalker/internal/time"
)

var LiveMatchModel = (*LiveMatch)(nil)

type LiveMatchID uint64

// LiveMatch ...
type LiveMatch struct {
	ID                         LiveMatchID        `gorm:"column:id;primary_key"`
	MatchID                    nspb.MatchID       `gorm:"column:match_id;unique_index;not null"`
	ServerSteamID              steamid.SteamId    `gorm:"column:server_steam_id;not null"`
	LobbyID                    nspb.LobbyID       `gorm:"column:lobby_id;unique_index;not null"`
	LobbyType                  nspb.LobbyType     `gorm:"column:lobby_type"`
	LeagueID                   nspb.LeagueID      `gorm:"column:league_id"`
	SeriesID                   nspb.SeriesID      `gorm:"column:series_id"`
	GameMode                   nspb.GameMode      `gorm:"column:game_mode"`
	AverageMMR                 uint32             `gorm:"column:average_mmr"`
	RadiantLead                int32              `gorm:"column:radiant_lead"`
	RadiantTeamID              TeamID             `gorm:"column:radiant_team_id"`
	RadiantTeamName            string             `gorm:"column:radiant_team_name;size:255"`
	RadiantTeamLogo            steamid.SteamId    `gorm:"column:radiant_team_logo"`
	RadiantScore               uint32             `gorm:"column:radiant_score"`
	DireTeamID                 TeamID             `gorm:"column:dire_team_id"`
	DireTeamName               string             `gorm:"column:dire_team_name;size:255"`
	DireTeamLogo               steamid.SteamId    `gorm:"column:dire_team_logo"`
	DireScore                  uint32             `gorm:"column:dire_score"`
	Delay                      uint32             `gorm:"column:delay"`
	ActivateTime               *time.Time         `gorm:"column:activate_time"`
	DeactivateTime             *time.Time         `gorm:"column:deactivate_time"`
	LastUpdateTime             *time.Time         `gorm:"column:last_update_time"`
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
}

func (*LiveMatch) TableName() string {
	return "live_matches"
}

func (m *LiveMatch) Equal(other *LiveMatch) bool {
	return m.MatchID == other.MatchID &&
		m.ServerSteamID == other.ServerSteamID &&
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
		nstime.EqualPtr(m.ActivateTime, other.ActivateTime) &&
		nstime.EqualPtr(m.DeactivateTime, other.DeactivateTime) &&
		nstime.EqualPtr(m.LastUpdateTime, other.LastUpdateTime) &&
		m.GameTime == other.GameTime &&
		m.Spectators == other.Spectators &&
		m.SortScore == other.SortScore &&
		m.BuildingState == other.BuildingState &&
		m.WeekendTourneyTournamentID == other.WeekendTourneyTournamentID &&
		m.WeekendTourneyDivision == other.WeekendTourneyDivision &&
		m.WeekendTourneySkillLevel == other.WeekendTourneySkillLevel &&
		m.WeekendTourneyBracketRound == other.WeekendTourneyBracketRound
}

func LiveMatchDotaProto(pb *protocol.CSourceTVGameSmall) *LiveMatch {
	return &LiveMatch{
		MatchID:                    pb.GetMatchId(),
		ServerSteamID:              steamid.SteamId(pb.GetServerSteamId()),
		LeagueID:                   nspb.LeagueID(pb.GetLeagueId()),
		SeriesID:                   nspb.SeriesID(pb.GetSeriesId()),
		LobbyID:                    pb.GetLobbyId(),
		LobbyType:                  nspb.LobbyType(pb.GetLobbyType()),
		GameMode:                   nspb.GameMode(pb.GetGameMode()),
		AverageMMR:                 pb.GetAverageMmr(),
		RadiantLead:                pb.GetRadiantLead(),
		RadiantTeamID:              TeamID(pb.GetTeamIdRadiant()),
		RadiantTeamName:            pb.GetTeamNameRadiant(),
		RadiantTeamLogo:            steamid.SteamId(TruncateUint(pb.GetTeamLogoRadiant())),
		RadiantScore:               pb.GetRadiantScore(),
		DireTeamID:                 TeamID(pb.GetTeamIdDire()),
		DireTeamName:               pb.GetTeamNameDire(),
		DireTeamLogo:               steamid.SteamId(TruncateUint(pb.GetTeamLogoDire())),
		DireScore:                  pb.GetDireScore(),
		ActivateTime:               NullUnixTimestamp(int64(pb.GetActivateTime())),
		DeactivateTime:             NullUnixTimestamp(int64(pb.GetDeactivateTime())),
		LastUpdateTime:             NullUnixTimestampFrac(float64(pb.GetLastUpdateTime())),
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
