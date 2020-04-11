package models

import (
	nssql "github.com/13k/night-stalker/internal/sql"
)

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
