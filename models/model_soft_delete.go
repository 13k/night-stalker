package models

import (
	"database/sql"
	"time"
)

type SoftDeletable interface {
	GetDeletedAt() time.Time
	SetDeletedAt(time.Time)
	IsDeleted() bool
}

var _ SoftDeletable = (*SoftDelete)(nil)

type SoftDelete struct {
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func (d SoftDelete) IsDeleted() bool {
	return d.DeletedAt.Valid
}

func (d SoftDelete) GetDeletedAt() time.Time {
	if d.DeletedAt.Valid {
		return d.DeletedAt.Time
	}

	return time.Time{}
}

func (d *SoftDelete) SetDeletedAt(t time.Time) {
	d.DeletedAt.Time, d.DeletedAt.Valid = t, true
}
