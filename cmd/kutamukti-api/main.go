package main

import (
	"kutamukti-api/pkg/config"
	"kutamukti-api/pkg/logger"
	"kutamukti-api/pkg/middleware"
	"kutamukti-api/routers"
	"log"
	"os"
	"time"

	internalRouters "kutamukti-api/internal/routers"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.InitEnvCheck()

	logger.Startup()
	port := os.Getenv("PORT")
	environment := os.Getenv("ENVIRONMENT")

	r := gin.New()
	r.Use(gin.Logger())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	db := config.InitDB()
	drive := config.InitDriveServie()
	validate := validator.New(validator.WithRequiredStructEnabled())
	lmt := tollbooth.NewLimiter(5, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Second})

	r.Use(middleware.GzipResponseMiddleware())
	r.Use(middleware.RateLimitMiddleware(lmt))

	internal := r.Group("/internal")
	internalRouters.InternalRouters(internal, db, validate)

	api := r.Group("/api")
	routers.CompRouters(api, db, validate, drive)

	var host string
	switch environment {
	case "development":
		host = "localhost"
	case "production":
		host = "0.0.0.0"
	default:
		panic("ENV ERROR: {ENVIRONMENT} UNKNOWN")
	}

	server := host + ":" + port
	err := r.Run(server)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}

	log.Println("Server started on port :" + port)
}
