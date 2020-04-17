package models

import (
	"database/sql"
)

var TeamTable = NewTable("teams")

type Team struct {
	ID `db:"id"`

	Name          string       `db:"name"`
	Tag           string       `db:"tag"`
	Rating        float32      `db:"rating"`
	Wins          uint32       `db:"wins"`
	Losses        uint32       `db:"losses"`
	LogoURL       string       `db:"logo_url"`
	LastMatchTime sql.NullTime `db:"last_match_time"`

	Timestamps
	SoftDelete

	Players []*Player `db:"-" model:"has_many"`
}
