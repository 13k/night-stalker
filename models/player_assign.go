// Code generated by modelgen. DO NOT EDIT.

package models

import (
	nssql "github.com/13k/night-stalker/internal/sql"
)

// Assign assigns fields from "other" into the receiver.
// It returns true if any changes were made to the receiver.
func (m *Player) Assign(other *Player) (dirty bool) {
	if other == nil {
		return false
	}

	if m.ID != other.ID {
		m.ID = other.ID
		dirty = true
	}

	if m.AccountID != other.AccountID {
		m.AccountID = other.AccountID
		dirty = true
	}

	if m.SteamID != other.SteamID {
		m.SteamID = other.SteamID
		dirty = true
	}

	if m.Name != other.Name {
		m.Name = other.Name
		dirty = true
	}

	if m.PersonaName != other.PersonaName {
		m.PersonaName = other.PersonaName
		dirty = true
	}

	if m.AvatarURL != other.AvatarURL {
		m.AvatarURL = other.AvatarURL
		dirty = true
	}

	if m.AvatarMediumURL != other.AvatarMediumURL {
		m.AvatarMediumURL = other.AvatarMediumURL
		dirty = true
	}

	if m.AvatarFullURL != other.AvatarFullURL {
		m.AvatarFullURL = other.AvatarFullURL
		dirty = true
	}

	if m.ProfileURL != other.ProfileURL {
		m.ProfileURL = other.ProfileURL
		dirty = true
	}

	if m.CountryCode != other.CountryCode {
		m.CountryCode = other.CountryCode
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
func (m *Player) AssignPartial(other *Player) (dirty bool) {
	if other == nil {
		return false
	}

	if other.ID != 0 && m.ID != other.ID {
		m.ID = other.ID
		dirty = true
	}

	if other.AccountID != 0 && m.AccountID != other.AccountID {
		m.AccountID = other.AccountID
		dirty = true
	}

	if other.SteamID != 0 && m.SteamID != other.SteamID {
		m.SteamID = other.SteamID
		dirty = true
	}

	if other.Name != "" && m.Name != other.Name {
		m.Name = other.Name
		dirty = true
	}

	if other.PersonaName != "" && m.PersonaName != other.PersonaName {
		m.PersonaName = other.PersonaName
		dirty = true
	}

	if other.AvatarURL != "" && m.AvatarURL != other.AvatarURL {
		m.AvatarURL = other.AvatarURL
		dirty = true
	}

	if other.AvatarMediumURL != "" && m.AvatarMediumURL != other.AvatarMediumURL {
		m.AvatarMediumURL = other.AvatarMediumURL
		dirty = true
	}

	if other.AvatarFullURL != "" && m.AvatarFullURL != other.AvatarFullURL {
		m.AvatarFullURL = other.AvatarFullURL
		dirty = true
	}

	if other.ProfileURL != "" && m.ProfileURL != other.ProfileURL {
		m.ProfileURL = other.ProfileURL
		dirty = true
	}

	if other.CountryCode != "" && m.CountryCode != other.CountryCode {
		m.CountryCode = other.CountryCode
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