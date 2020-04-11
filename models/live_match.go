package models

import (
	"database/sql"

	d2pb "github.com/paralin/go-dota2/protocol"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nssql "github.com/13k/night-stalker/internal/sql"
)

var LiveMatchTable = NewTable("live_matches")

type LiveMatch struct {
	ID `db:"id" goqu:"defaultifempty"`

	ServerID                   nspb.SteamID       `db:"server_id"`
	LobbyID                    nspb.LobbyID       `db:"lobby_id"`
	LobbyType                  nspb.LobbyType     `db:"lobby_type"`
	SeriesID                   nspb.SeriesID      `db:"series_id"`
	GameMode                   nspb.GameMode      `db:"game_mode"`
	AverageMMR                 uint32             `db:"average_mmr"`
	RadiantLead                int32              `db:"radiant_lead"`
	RadiantTeamName            string             `db:"radiant_team_name"`
	RadiantTeamLogo            nspb.SteamID       `db:"radiant_team_logo"`
	RadiantScore               uint32             `db:"radiant_score"`
	DireTeamName               string             `db:"dire_team_name"`
	DireTeamLogo               nspb.SteamID       `db:"dire_team_logo"`
	DireScore                  uint32             `db:"dire_score"`
	Delay                      uint32             `db:"delay"`
	ActivateTime               sql.NullTime       `db:"activate_time"`
	DeactivateTime             sql.NullTime       `db:"deactivate_time"`
	LastUpdateTime             sql.NullTime       `db:"last_update_time"`
	GameTime                   int32              `db:"game_time"`
	Spectators                 uint32             `db:"spectators"`
	SortScore                  float64            `db:"sort_score"`
	BuildingState              nspb.BuildingState `db:"building_state"`
	WeekendTourneyTournamentID uint32             `db:"weekend_tourney_tournament_id"`
	WeekendTourneyDivision     uint32             `db:"weekend_tourney_division"`
	WeekendTourneySkillLevel   uint32             `db:"weekend_tourney_skill_level"`
	WeekendTourneyBracketRound uint32             `db:"weekend_tourney_bracket_round"`

	MatchID       ID `db:"match_id"`
	LeagueID      ID `db:"league_id"`
	RadiantTeamID ID `db:"radiant_team_id"`
	DireTeamID    ID `db:"dire_team_id"`

	Timestamps
	SoftDelete

	Match       *Match             `db:"-" model:"belongs_to"`
	League      *League            `db:"-" model:"belongs_to"`
	RadiantTeam *Team              `db:"-" model:"belongs_to"`
	DireTeam    *Team              `db:"-" model:"belongs_to"`
	Players     []*LiveMatchPlayer `db:"-" model:"has_many"`
	Stats       []*LiveMatchStats  `db:"-" model:"has_many"`
}

func NewLiveMatchProto(pb *d2pb.CSourceTVGameSmall) *LiveMatch {
	return &LiveMatch{
		MatchID:                    ID(pb.GetMatchId()),
		ServerID:                   nspb.SteamID(pb.GetServerSteamId()),
		LeagueID:                   ID(pb.GetLeagueId()),
		SeriesID:                   nspb.SeriesID(pb.GetSeriesId()),
		LobbyID:                    nspb.LobbyID(pb.GetLobbyId()),
		LobbyType:                  nspb.LobbyType(pb.GetLobbyType()),
		GameMode:                   nspb.GameMode(pb.GetGameMode()),
		AverageMMR:                 pb.GetAverageMmr(),
		RadiantLead:                pb.GetRadiantLead(),
		RadiantTeamID:              ID(pb.GetTeamIdRadiant()),
		RadiantTeamName:            pb.GetTeamNameRadiant(),
		RadiantTeamLogo:            nspb.SteamID(TruncateUint(pb.GetTeamLogoRadiant())),
		RadiantScore:               pb.GetRadiantScore(),
		DireTeamID:                 ID(pb.GetTeamIdDire()),
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
