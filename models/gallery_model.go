package models

import "time"

type ItemTypes string

const (
	Image ItemTypes = "image"
	Video ItemTypes = "video"
)

type Galleries struct {
	ID          int64     `gorm:"primaryKey"`
	UUID        string    `gorm:"not null;unique;index"`
	Type        ItemTypes `gorm:"not null;index"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Slug        string    `gorm:"not null;unique;index"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
