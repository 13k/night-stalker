package models

import (
	"database/sql"
	"time"
)

type Timestamps struct {
	CreatedAt time.Time    `gorm:"column:created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at"`
	DeletedAt sql.NullTime `gorm:"column:deleted_at;index"`
}
