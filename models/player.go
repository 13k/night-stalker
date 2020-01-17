package models

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
	"github.com/faceit/go-steam/steamid"
)

var PlayerModel = (*Player)(nil)

type PlayerID uint64

// Player ...
type Player struct {
	ID              PlayerID        `gorm:"column:id;primary_key"`
	AccountID       nspb.AccountID  `gorm:"column:account_id;unique_index;not null"`
	SteamID         steamid.SteamId `gorm:"column:steam_id"`
	Name            string          `gorm:"column:name"`
	PersonaName     string          `gorm:"column:persona_name"`
	AvatarURL       string          `gorm:"column:avatar_url"`
	AvatarMediumURL string          `gorm:"column:avatar_medium_url"`
	AvatarFullURL   string          `gorm:"column:avatar_full_url"`
	ProfileURL      string          `gorm:"column:profile_url"`
	CountryCode     string          `gorm:"column:country_code"`
	Timestamps
}

func (*Player) TableName() string {
	return "players"
}
