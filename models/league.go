package models

import (
	"database/sql"

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
	LastActivityAt sql.NullTime      `gorm:"column:last_activity_at"`
	StartAt        sql.NullTime      `gorm:"column:start_at"`
	FinishAt       sql.NullTime      `gorm:"column:finish_at"`
	Timestamps
}

func (*League) TableName() string {
	return "leagues"
}
