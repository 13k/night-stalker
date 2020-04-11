package models

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

var PlayerTable = NewTable("players")

type Player struct {
	ID `db:"id" goqu:"defaultifempty"`

	AccountID       nspb.AccountID `db:"account_id"`
	SteamID         nspb.SteamID   `db:"steam_id"`
	Name            string         `db:"name"`
	PersonaName     string         `db:"persona_name"`
	AvatarURL       string         `db:"avatar_url"`
	AvatarMediumURL string         `db:"avatar_medium_url"`
	AvatarFullURL   string         `db:"avatar_full_url"`
	ProfileURL      string         `db:"profile_url"`
	CountryCode     string         `db:"country_code"`

	Timestamps
	SoftDelete
}
