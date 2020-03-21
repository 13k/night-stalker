package models

import (
	"time"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

var ProPlayerModel Model = (*ProPlayer)(nil)

type ProPlayerID uint64

// ProPlayer ...
type ProPlayer struct {
	ID          ProPlayerID      `gorm:"column:id;primary_key"`
	AccountID   nspb.AccountID   `gorm:"column:account_id;unique_index;not null"`
	TeamID      nspb.TeamID      `gorm:"column:team_id"`
	IsLocked    bool             `gorm:"column:is_locked"`
	LockedUntil *time.Time       `gorm:"column:locked_until"`
	FantasyRole nspb.FantasyRole `gorm:"column:fantasy_role"`
	Timestamps

	Team *Team
}

func (*ProPlayer) TableName() string {
	return "pro_players"
}
