// Code generated by modelgen. DO NOT EDIT.

package models

import (
	nssql "github.com/13k/night-stalker/internal/sql"
)

// Assign assigns fields from "other" into the receiver.
// It returns true if any changes were made to the receiver.
func (m *Match) Assign(other *Match) (dirty bool) {
	if other == nil {
		return false
	}

	if m.ID != other.ID {
		m.ID = other.ID
		dirty = true
	}

	if m.SeriesType != other.SeriesType {
		m.SeriesType = other.SeriesType
		dirty = true
	}

	if m.SeriesGame != other.SeriesGame {
		m.SeriesGame = other.SeriesGame
		dirty = true
	}

	if m.GameMode != other.GameMode {
		m.GameMode = other.GameMode
		dirty = true
	}

	if !nssql.NullTimeEqual(m.StartTime, other.StartTime) {
		m.StartTime = other.StartTime
		dirty = true
	}

	if m.Duration != other.Duration {
		m.Duration = other.Duration
		dirty = true
	}

	if m.Outcome != other.Outcome {
		m.Outcome = other.Outcome
		dirty = true
	}

	if m.RadiantTeamName != other.RadiantTeamName {
		m.RadiantTeamName = other.RadiantTeamName
		dirty = true
	}

	if m.RadiantTeamLogo != other.RadiantTeamLogo {
		m.RadiantTeamLogo = other.RadiantTeamLogo
		dirty = true
	}

	if m.RadiantTeamLogoURL != other.RadiantTeamLogoURL {
		m.RadiantTeamLogoURL = other.RadiantTeamLogoURL
		dirty = true
	}

	if m.RadiantScore != other.RadiantScore {
		m.RadiantScore = other.RadiantScore
		dirty = true
	}

	if m.DireTeamName != other.DireTeamName {
		m.DireTeamName = other.DireTeamName
		dirty = true
	}

	if m.DireTeamLogo != other.DireTeamLogo {
		m.DireTeamLogo = other.DireTeamLogo
		dirty = true
	}

	if m.DireTeamLogoURL != other.DireTeamLogoURL {
		m.DireTeamLogoURL = other.DireTeamLogoURL
		dirty = true
	}

	if m.DireScore != other.DireScore {
		m.DireScore = other.DireScore
		dirty = true
	}

	if m.WeekendTourneyTournamentID != other.WeekendTourneyTournamentID {
		m.WeekendTourneyTournamentID = other.WeekendTourneyTournamentID
		dirty = true
	}

	if m.WeekendTourneySeasonTrophyID != other.WeekendTourneySeasonTrophyID {
		m.WeekendTourneySeasonTrophyID = other.WeekendTourneySeasonTrophyID
		dirty = true
	}

	if m.WeekendTourneyDivision != other.WeekendTourneyDivision {
		m.WeekendTourneyDivision = other.WeekendTourneyDivision
		dirty = true
	}

	if m.WeekendTourneySkillLevel != other.WeekendTourneySkillLevel {
		m.WeekendTourneySkillLevel = other.WeekendTourneySkillLevel
		dirty = true
	}

	if m.LeagueID != other.LeagueID {
		m.LeagueID = other.LeagueID
		dirty = true
	}

	if m.RadiantTeamID != other.RadiantTeamID {
		m.RadiantTeamID = other.RadiantTeamID
		dirty = true
	}

	if m.DireTeamID != other.DireTeamID {
		m.DireTeamID = other.DireTeamID
		dirty = true
	}

	if !m.CreatedAt.Equal(other.CreatedAt) {
		m.CreatedAt = other.CreatedAt
		dirty = true
	}

	if !m.UpdatedAt.Equal(other.UpdatedAt) {
		m.UpdatedAt = other.UpdatedAt
		dirty = true
	}

	if !nssql.NullTimeEqual(m.DeletedAt, other.DeletedAt) {
		m.DeletedAt = other.DeletedAt
		dirty = true
	}

	return
}

// AssignPartial assigns fields with non-zero values from "other" into the receiver.
// It returns true if any changes were made to the receiver.
func (m *Match) AssignPartial(other *Match) (dirty bool) {
	if other == nil {
		return false
	}

	if other.ID != 0 && m.ID != other.ID {
		m.ID = other.ID
		dirty = true
	}

	if other.SeriesType != 0 && m.SeriesType != other.SeriesType {
		m.SeriesType = other.SeriesType
		dirty = true
	}

	if other.SeriesGame != 0 && m.SeriesGame != other.SeriesGame {
		m.SeriesGame = other.SeriesGame
		dirty = true
	}

	if other.GameMode != 0 && m.GameMode != other.GameMode {
		m.GameMode = other.GameMode
		dirty = true
	}

	if !nssql.NullTimeIsZero(other.StartTime) && !nssql.NullTimeEqual(m.StartTime, other.StartTime) {
		m.StartTime = other.StartTime
		dirty = true
	}

	if other.Duration != 0 && m.Duration != other.Duration {
		m.Duration = other.Duration
		dirty = true
	}

	if other.Outcome != 0 && m.Outcome != other.Outcome {
		m.Outcome = other.Outcome
		dirty = true
	}

	if other.RadiantTeamName != "" && m.RadiantTeamName != other.RadiantTeamName {
		m.RadiantTeamName = other.RadiantTeamName
		dirty = true
	}

	if other.RadiantTeamLogo != 0 && m.RadiantTeamLogo != other.RadiantTeamLogo {
		m.RadiantTeamLogo = other.RadiantTeamLogo
		dirty = true
	}

	if other.RadiantTeamLogoURL != "" && m.RadiantTeamLogoURL != other.RadiantTeamLogoURL {
		m.RadiantTeamLogoURL = other.RadiantTeamLogoURL
		dirty = true
	}

	if other.RadiantScore != 0 && m.RadiantScore != other.RadiantScore {
		m.RadiantScore = other.RadiantScore
		dirty = true
	}

	if other.DireTeamName != "" && m.DireTeamName != other.DireTeamName {
		m.DireTeamName = other.DireTeamName
		dirty = true
	}

	if other.DireTeamLogo != 0 && m.DireTeamLogo != other.DireTeamLogo {
		m.DireTeamLogo = other.DireTeamLogo
		dirty = true
	}

	if other.DireTeamLogoURL != "" && m.DireTeamLogoURL != other.DireTeamLogoURL {
		m.DireTeamLogoURL = other.DireTeamLogoURL
		dirty = true
	}

	if other.DireScore != 0 && m.DireScore != other.DireScore {
		m.DireScore = other.DireScore
		dirty = true
	}

	if other.WeekendTourneyTournamentID != 0 && m.WeekendTourneyTournamentID != other.WeekendTourneyTournamentID {
		m.WeekendTourneyTournamentID = other.WeekendTourneyTournamentID
		dirty = true
	}

	if other.WeekendTourneySeasonTrophyID != 0 && m.WeekendTourneySeasonTrophyID != other.WeekendTourneySeasonTrophyID {
		m.WeekendTourneySeasonTrophyID = other.WeekendTourneySeasonTrophyID
		dirty = true
	}

	if other.WeekendTourneyDivision != 0 && m.WeekendTourneyDivision != other.WeekendTourneyDivision {
		m.WeekendTourneyDivision = other.WeekendTourneyDivision
		dirty = true
	}

	if other.WeekendTourneySkillLevel != 0 && m.WeekendTourneySkillLevel != other.WeekendTourneySkillLevel {
		m.WeekendTourneySkillLevel = other.WeekendTourneySkillLevel
		dirty = true
	}

	if other.LeagueID != 0 && m.LeagueID != other.LeagueID {
		m.LeagueID = other.LeagueID
		dirty = true
	}

	if other.RadiantTeamID != 0 && m.RadiantTeamID != other.RadiantTeamID {
		m.RadiantTeamID = other.RadiantTeamID
		dirty = true
	}

	if other.DireTeamID != 0 && m.DireTeamID != other.DireTeamID {
		m.DireTeamID = other.DireTeamID
		dirty = true
	}

	if !other.CreatedAt.IsZero() && !m.CreatedAt.Equal(other.CreatedAt) {
		m.CreatedAt = other.CreatedAt
		dirty = true
	}

	if !other.UpdatedAt.IsZero() && !m.UpdatedAt.Equal(other.UpdatedAt) {
		m.UpdatedAt = other.UpdatedAt
		dirty = true
	}

	if !nssql.NullTimeIsZero(other.DeletedAt) && !nssql.NullTimeEqual(m.DeletedAt, other.DeletedAt) {
		m.DeletedAt = other.DeletedAt
		dirty = true
	}

	return
}