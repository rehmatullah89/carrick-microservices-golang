package db

import (
	"database/sql"
	"time"
)

type Timestamps struct {
	CreatedAt    time.Time    `gorm:"default:current_timestamp"`
	UpdatedAt    time.Time    `gorm:"default:current_timestamp"`
	DeletedAt    sql.NullTime `gorm:"index"`
}
