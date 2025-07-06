package routers

import (
	"kutamukti-api/injectors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CompRouters(r *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "pong",
		})
	})

	userController := injectors.InitializeUserController(db, validate)
	complaintController := injectors.InitializeComplaintController(db, validate)
	scheduleController := injectors.InitializeScheduleController(db, validate)
	announcementController := injectors.InitializeAnnouncementController(db, validate)
	umkmController := injectors.InitializeUMKMController(db, validate)
	galleryController := injectors.InitializeGalleryController(db, validate)
	storageController := injectors.InitializeStorageController(db, validate)

	UserRoutes(r, userController)
	ComplaintRoutes(r, complaintController)
	ScheduleRoutes(r, scheduleController)
	AnnouncementRoutes(r, announcementController)
	UMKMRoutes(r, umkmController)
	GalleryRoutes(r, galleryController)
	StorageRoutes(r, storageController)
}
