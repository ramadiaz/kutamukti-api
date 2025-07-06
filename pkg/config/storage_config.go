package config

import (
	"log"
	"os"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func InitDriveServie() *drive.Service {
	APPLICATION_CREDENTIALS := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	driveService, err := drive.NewService(nil, option.WithCredentialsJSON([]byte(APPLICATION_CREDENTIALS)))
	if err != nil {
		log.Fatal(err)
	}

	return driveService
}
