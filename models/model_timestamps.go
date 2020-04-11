package models

import (
	"time"
)

type Timestampable interface {
	GetCreatedAt() time.Time
	SetCreatedAt(time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(time.Time)
}

var _ Timestampable = (*Timestamps)(nil)

type Timestamps struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (ts Timestamps) GetCreatedAt() time.Time   { return ts.CreatedAt }
func (ts *Timestamps) SetCreatedAt(t time.Time) { ts.CreatedAt = t }

func (ts Timestamps) GetUpdatedAt() time.Time   { return ts.UpdatedAt }
func (ts *Timestamps) SetUpdatedAt(t time.Time) { ts.UpdatedAt = t }
