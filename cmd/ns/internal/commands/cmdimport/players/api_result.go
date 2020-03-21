package players

import (
	nsjson "github.com/13k/night-stalker/internal/json"
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

type apiResult []*apiResultEntry

type apiResultEntry struct {
	AccountID    nspb.AccountID    `json:"account_id,omitempty"`
	SteamID      nsjson.StringUint `json:"steamid,omitempty"`
	TeamID       nspb.TeamID       `json:"team_id,omitempty"`
	Name         string            `json:"name,omitempty"`
	PersonaName  string            `json:"personaname,omitempty"`
	Avatar       string            `json:"avatar,omitempty"`
	AvatarMedium string            `json:"avatarmedium,omitempty"`
	AvatarFull   string            `json:"avatarfull,omitempty"`
	ProfileURL   string            `json:"profileurl,omitempty"`
	CountryCode  string            `json:"loccountrycode,omitempty"`
	FantasyRole  nspb.FantasyRole  `json:"fantasy_role,omitempty"`
	IsLocked     bool              `json:"is_locked,omitempty"`
	LockedUntil  *nsjson.UnixTime  `json:"locked_until,omitempty"`
}
