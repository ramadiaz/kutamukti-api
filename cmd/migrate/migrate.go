package main

import (
	"kutamukti-api/pkg/config"
	"kutamukti-api/models"
)

func main() {
	db := config.InitDB()

	err := db.AutoMigrate(
		&models.Users{},
	)
	if err != nil {
		panic("failed to migrate models: " + err.Error())
	}
}
