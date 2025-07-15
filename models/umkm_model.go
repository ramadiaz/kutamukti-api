package models

import (
	"time"

	"gorm.io/gorm"
)

type UMKM struct {
	gorm.Model

	ID          int64   `gorm:"primaryKey"`
	UUID        string  `gorm:"not null;unique;index"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Is24Hours   bool    `gorm:"not null;default:false"`
	Location    string  `gorm:"not null"`
	Langitude   float64 `gorm:"not null"`
	Latitude    float64 `gorm:"not null"`
	OpenAt      int     `gorm:"not null"`
	CloseAt     int     `gorm:"not null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`

	Products []UMKMProduct `gorm:"foreignKey:UMKMUUID;references:UUID;constraint:OnDelete:CASCADE"`
}

type UMKMProduct struct {
	gorm.Model
	ID          int64      `gorm:"primaryKey"`
	UUID        string     `gorm:"not null;unique;index"`
	UMKMUUID    string     `gorm:"not null;index"`
	Name        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	Variation   string     `gorm:"not null"`
	Price       int        `gorm:"not null"`
	CreatedAt   time.Time  `gorm:"not null"`
	UpdatedAt   time.Time  `gorm:"not null"`
	DeletedAt   *time.Time `gorm:"index"`

	Images []UMKMImage `gorm:"foreignKey:UMKMProductUUID;references:UUID;constraint:OnDelete:CASCADE"`
}

type UMKMImage struct {
	gorm.Model
	ID              int64      `gorm:"primaryKey"`
	UMKMProductUUID string     `gorm:"not null;index"`
	ImageURL        string     `gorm:"not null"`
	CreatedAt       time.Time  `gorm:"not null"`
	UpdatedAt       time.Time  `gorm:"not null"`
	DeletedAt       *time.Time `gorm:"index"`
}
