// Code generated by modelgen. DO NOT EDIT.

package models

import (
	nssql "github.com/13k/night-stalker/internal/sql"
)

// Assign assigns fields from "other" into the receiver.
// It returns true if any changes were made to the receiver.
func (m *Hero) Assign(other *Hero) (dirty bool) {
	if other == nil {
		return false
	}

	if m.ID != other.ID {
		m.ID = other.ID
		dirty = true
	}

	if m.Name != other.Name {
		m.Name = other.Name
		dirty = true
	}

	if m.Slug != other.Slug {
		m.Slug = other.Slug
		dirty = true
	}

	if m.LocalizedName != other.LocalizedName {
		m.LocalizedName = other.LocalizedName
		dirty = true
	}

	if !StringsEqual(m.Aliases, other.Aliases) {
		m.Aliases = other.Aliases
		dirty = true
	}

	if !nssql.IntArrayEqual(m.Roles, other.Roles) {
		m.Roles = other.Roles
		dirty = true
	}

	if !IntsEqual(m.RoleLevels, other.RoleLevels) {
		m.RoleLevels = other.RoleLevels
		dirty = true
	}

	if m.Complexity != other.Complexity {
		m.Complexity = other.Complexity
		dirty = true
	}

	if m.Legs != other.Legs {
		m.Legs = other.Legs
		dirty = true
	}

	if m.AttributePrimary != other.AttributePrimary {
		m.AttributePrimary = other.AttributePrimary
		dirty = true
	}

	if m.AttackCapabilities != other.AttackCapabilities {
		m.AttackCapabilities = other.AttackCapabilities
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
func (m *Hero) AssignPartial(other *Hero) (dirty bool) {
	if other == nil {
		return false
	}

	if other.ID != 0 && m.ID != other.ID {
		m.ID = other.ID
		dirty = true
	}

	if other.Name != "" && m.Name != other.Name {
		m.Name = other.Name
		dirty = true
	}

	if other.Slug != "" && m.Slug != other.Slug {
		m.Slug = other.Slug
		dirty = true
	}

	if other.LocalizedName != "" && m.LocalizedName != other.LocalizedName {
		m.LocalizedName = other.LocalizedName
		dirty = true
	}

	if other.Aliases != nil && !StringsEqual(m.Aliases, other.Aliases) {
		m.Aliases = other.Aliases
		dirty = true
	}

	if other.Roles != nil && !nssql.IntArrayEqual(m.Roles, other.Roles) {
		m.Roles = other.Roles
		dirty = true
	}

	if other.RoleLevels != nil && !IntsEqual(m.RoleLevels, other.RoleLevels) {
		m.RoleLevels = other.RoleLevels
		dirty = true
	}

	if other.Complexity != 0 && m.Complexity != other.Complexity {
		m.Complexity = other.Complexity
		dirty = true
	}

	if other.Legs != 0 && m.Legs != other.Legs {
		m.Legs = other.Legs
		dirty = true
	}

	if other.AttributePrimary != 0 && m.AttributePrimary != other.AttributePrimary {
		m.AttributePrimary = other.AttributePrimary
		dirty = true
	}

	if other.AttackCapabilities != 0 && m.AttackCapabilities != other.AttackCapabilities {
		m.AttackCapabilities = other.AttackCapabilities
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
