package models

import (
	"database/sql"

	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
)

var ProPlayerTable = NewTable("pro_players")

type ProPlayer struct {
	ID `db:"id" goqu:"defaultifempty"`

	AccountID   nspb.AccountID   `db:"account_id"`
	IsLocked    bool             `db:"is_locked"`
	LockedUntil sql.NullTime     `db:"locked_until"`
	FantasyRole nspb.FantasyRole `db:"fantasy_role"`

	TeamID ID `db:"team_id"`

	Timestamps
	SoftDelete

	Team *Team `db:"-" model:"belongs_to"`
}
