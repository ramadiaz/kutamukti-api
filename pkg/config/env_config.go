package config

import (
	"kutamukti-api/pkg/logger"
	"os"
	"reflect"
)

type Env struct {
	DB_HOST                      string
	DB_NAME                      string
	DB_PASSWORD                  string
	DB_PORT                      string
	DB_USER                      string
	PORT                         string
	JWT_SECRET                   string
	ADMIN_PASSWORD               string
	ADMIN_USERNAME               string
	FONNTE_API_KEY               string
	FONNTE_GROUP_ANNOUNCEMENT_ID string
	FONNTE_GROUP_COMPLAINT_ID    string
	ENVIRONMENT                  string
	SMTP_EMAIL                   string
	SMTP_PASSWORD                string
	SMTP_SERVER                  string
	SMTP_PORT                    string
}

func InitEnvCheck() {
	logger.Info("Checking environment variables...")
	environment := Env{
		DB_HOST:                      os.Getenv("DB_HOST"),
		DB_NAME:                      os.Getenv("DB_NAME"),
		DB_PASSWORD:                  os.Getenv("DB_PASSWORD"),
		DB_PORT:                      os.Getenv("DB_PORT"),
		DB_USER:                      os.Getenv("DB_USER"),
		PORT:                         os.Getenv("PORT"),
		JWT_SECRET:                   os.Getenv("JWT_SECRET"),
		ADMIN_PASSWORD:               os.Getenv("ADMIN_PASSWORD"),
		ADMIN_USERNAME:               os.Getenv("ADMIN_USERNAME"),
		FONNTE_API_KEY:               os.Getenv("FONNTE_API_KEY"),
		FONNTE_GROUP_ANNOUNCEMENT_ID: os.Getenv("FONNTE_GROUP_ANNOUNCEMENT_ID"),
		FONNTE_GROUP_COMPLAINT_ID:    os.Getenv("FONNTE_GROUP_COMPLAINT_ID"),
		ENVIRONMENT:                  os.Getenv("ENVIRONMENT"),
		SMTP_EMAIL:                   os.Getenv("SMTP_EMAIL"),
		SMTP_PASSWORD:                os.Getenv("SMTP_PASSWORD"),
		SMTP_SERVER:                  os.Getenv("SMTP_SERVER"),
		SMTP_PORT:                    os.Getenv("SMTP_PORT"),
	}

	isEmpty, emptyFields := checkEmptyFields(environment)
	if isEmpty {
		logger.PanicError("Missing environment variables: %v", emptyFields)
	} else {
		logger.Info("Environment variables are set!")
	}
}

func checkEmptyFields(env Env) (bool, []string) {
	v := reflect.ValueOf(env)
	typeOfEnv := v.Type()
	var emptyFields []string

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String && field.String() == "" {
			emptyFields = append(emptyFields, typeOfEnv.Field(i).Name)
		}
	}

	return len(emptyFields) > 0, emptyFields
}
