package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedules struct {
	gorm.Model

	ID          uint64 `gorm:"primaryKey"`
	UUID        string `gorm:"unique;index"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Location    string `gorm:"not null"`

	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}
