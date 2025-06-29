package models

import (
	"time"

	"gorm.io/gorm"
)

type Roles string

const (
	Admin Roles = "admin"
	Staff Roles = "staff"
)

type Users struct {
	gorm.Model

	ID             int64  `gorm:"primaryKey"`
	UUID           string `gorm:"not null;unique;index"`
	Email          string `gorm:"not null;unique;index"`
	HashedPassword string `gorm:"not null"`
	Username       string `gorm:"not null;unique;index"`
	Name           string `gorm:"not null"`
	Role           Roles  `gorm:"not null"`

	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}
