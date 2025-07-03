package models

import (
	"time"

	"gorm.io/gorm"
)

type ComplaintStatus string

const (
	Open    ComplaintStatus = "open"
	Process ComplaintStatus = "process"
	Closed  ComplaintStatus = "closed"
)

type Complaints struct {
	gorm.Model

	ID          int64           `gorm:"primaryKey"`
	UUID        string          `gorm:"not null;unique;index"`
	Title       string          `gorm:"not null"`
	Description string          `gorm:"not null"`
	Status      ComplaintStatus `gorm:"not null;default:open"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}
