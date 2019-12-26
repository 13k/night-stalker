package models

import (
	"time"

	nsproto "github.com/13k/night-stalker/internal/protocol"
)

var ProPlayerModel = (*ProPlayer)(nil)

type ProPlayerID uint64

// ProPlayer ...
type ProPlayer struct {
	ID          ProPlayerID         `gorm:"column:id;primary_key"`
	AccountID   nsproto.AccountID   `gorm:"column:account_id;unique_index;not null"`
	TeamID      TeamID              `gorm:"column:team_id"`
	IsLocked    bool                `gorm:"column:is_locked"`
	LockedUntil *time.Time          `gorm:"column:locked_until"`
	FantasyRole nsproto.FantasyRole `gorm:"column:fantasy_role"`
	Timestamps

	Team *Team
}
