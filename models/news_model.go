package models

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model

	ID           int64  `gorm:"primaryKey"`
	UUID         string `gorm:"not null;unique;index"`
	Title        string `gorm:"not null"`
	Content      string `gorm:"not null"`
	RawText      string `gorm:"not null"`
	Slug         string `gorm:"not null;unique;index"`
	ThumbnailURL string `gorm:"not null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`

	Images []NewsImages `gorm:"foreignKey:NewsID"`
}

type NewsImages struct {
	gorm.Model

	ID       int64  `gorm:"primaryKey"`
	NewsID   string `gorm:"not null;index"`
	ImageURL string `gorm:"not null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}
