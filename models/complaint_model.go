package models

import "time"

type Complaints struct {
	ID          int64  `gorm:"primaryKey"`
	UUID        string `gorm:"unique;index"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}