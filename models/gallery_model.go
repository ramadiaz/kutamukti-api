package models

import (
	"time"

	"gorm.io/gorm"
)

type ImageGalleries struct {
	ID    int64  `gorm:"primaryKey"`
	UUID  string `gorm:"not null;unique;index"`
	Title string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`

	Images []Images `gorm:"foreignKey:GalleryUUID;references:UUID;constraint:OnDelete:CASCADE"`
}

func (u *ImageGalleries) BeforeDelete(tx *gorm.DB) (err error) {
	return tx.Model(&Images{}).Where("gallery_uuid = ?", u.UUID).Delete(&Images{}).Error
}

type Images struct {
	ID          int64  `gorm:"primaryKey"`
	GalleryUUID string `gorm:"not null;index"`
	ImageURL    string `gorm:"not null"`
	Description string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

type Videos struct {
	ID          int64  `gorm:"primaryKey"`
	UUID        string `gorm:"not null;unique;index"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	YoutubeURL  string `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
