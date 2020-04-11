package models

import (
	nspb "github.com/13k/night-stalker/internal/protobuf/protocol"
	nsstr "github.com/13k/night-stalker/internal/strings"
)

var FollowedPlayerTable = NewTable("followed_players")

type FollowedPlayer struct {
	ID `db:"id" goqu:"defaultifempty"`

	AccountID nspb.AccountID `db:"account_id"`
	Label     string         `db:"label"`
	Slug      string         `db:"slug"`

	Timestamps
	SoftDelete
}

func (m *FollowedPlayer) BeforeCreate() error {
	if m.Slug == "" {
		m.Slug = nsstr.Slugify(m.Label)
	}

	return nil
}
