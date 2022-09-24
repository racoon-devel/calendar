package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// Meeting описывает одиночную или повторяющуюся встречу
type Meeting struct {
	gorm.Model
	Owner       uint   `gorm:"not null"`
	User        User   `gorm:"foreignKey:owner"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Private     bool   `gorm:"not null"`
	Notify      sql.NullInt32
	RRule       sql.NullString
	StartTime   time.Time     `gorm:"not null"`
	Duration    time.Duration `gorm:"not null"`
}
