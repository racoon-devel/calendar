package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// Meeting описывает одиночную или повторяющуюся встречу
type Meeting struct {
	gorm.Model
	Title       string
	Description string
	Private     bool
	Notify      sql.NullInt32
	RRule       string
	StartTime   time.Time
	DurationMin uint
}
