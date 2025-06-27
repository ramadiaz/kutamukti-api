package models

import (
	"time"

	"gorm.io/gorm"
)

type Announcements struct {
	gorm.Model

	ID          int64  `gorm:"primaryKey"`
	UUID        string `gorm:"not null;unique;index"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Slug        string `gorm:"not null;unique;index"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}
