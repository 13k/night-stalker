package models

import (
	"time"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

var LeagueModel Model = (*League)(nil)

type League struct {
	ID             nspb.LeagueID     `gorm:"column:id;primary_key"`
	Name           string            `gorm:"column:name;size:255;not null"`
	Tier           nspb.LeagueTier   `gorm:"column:tier;not null"`
	Region         nspb.LeagueRegion `gorm:"column:region;not null"`
	Status         nspb.LeagueStatus `gorm:"column:status;not null"`
	TotalPrizePool uint32            `gorm:"column:total_prize_pool"`
	LastActivityAt *time.Time        `gorm:"column:last_activity_at"`
	StartAt        *time.Time        `gorm:"column:start_at"`
	FinishAt       *time.Time        `gorm:"column:finish_at"`
	Timestamps
}

func (*League) TableName() string {
	return "leagues"
}
