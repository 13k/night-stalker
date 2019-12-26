package models

import (
	nsproto "github.com/13k/night-stalker/internal/protocol"
)

var FollowedPlayerModel = (*FollowedPlayer)(nil)

type FollowedPlayerID uint64

// FollowedPlayer ...
type FollowedPlayer struct {
	ID        FollowedPlayerID  `gorm:"column:id;primary_key"`
	AccountID nsproto.AccountID `gorm:"column:account_id;unique_index;not null"`
	Label     string            `gorm:"column:label;size:255;not null"`
	Timestamps
}
