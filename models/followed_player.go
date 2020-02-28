package models

import (
	nspb "github.com/13k/night-stalker/internal/protocol"
)

var FollowedPlayerModel Model = (*FollowedPlayer)(nil)

type FollowedPlayerID uint64

// FollowedPlayer ...
type FollowedPlayer struct {
	ID        FollowedPlayerID `gorm:"column:id;primary_key"`
	AccountID nspb.AccountID   `gorm:"column:account_id;unique_index;not null"`
	Label     string           `gorm:"column:label;size:255;not null"`
	Timestamps
}

func (*FollowedPlayer) TableName() string {
	return "followed_players"
}
