package models

import (
	"database/sql"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

var TeamModel Model = (*Team)(nil)

// Team ...
type Team struct {
	ID            nspb.TeamID  `gorm:"column:id;primary_key"`
	Name          string       `gorm:"column:name;size:255;not null"`
	Tag           string       `gorm:"column:tag;size:255;not null"`
	Rating        float32      `gorm:"column:rating"`
	Wins          uint32       `gorm:"column:wins"`
	Losses        uint32       `gorm:"column:losses"`
	LogoURL       string       `gorm:"column:logo_url"`
	LastMatchTime sql.NullTime `gorm:"column:last_match_time"`
	Timestamps
}

func (*Team) TableName() string {
	return "teams"
}
