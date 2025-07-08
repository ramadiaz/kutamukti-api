package main

import (
	"kutamukti-api/models"
	"kutamukti-api/pkg/config"
)

func main() {
	db := config.InitDB()

	err := db.AutoMigrate(
		&models.Complaints{},
		&models.Schedules{},
		&models.Announcements{},
		&models.UMKM{},
		&models.Galleries{},
		&models.Users{},
		&models.Files{},
		&models.News{},
		&models.NewsImages{},
	)
	if err != nil {
		panic("failed to migrate models: " + err.Error())
	}
}
