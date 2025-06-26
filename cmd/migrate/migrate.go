package main

import (
	"kutamukti-api/models"
	"kutamukti-api/pkg/config"
)

func main() {
	db := config.InitDB()

	err := db.AutoMigrate(
		&models.Complaints{},
	)
	if err != nil {
		panic("failed to migrate models: " + err.Error())
	}
}
