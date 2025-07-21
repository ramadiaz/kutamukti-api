package routers

import (
	"kutamukti-api/api/gallery/controllers"
	"kutamukti-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func GalleryRoutes(r *gin.RouterGroup, galleryController controllers.CompControllers) {
	galleryGroup := r.Group("/gallery")
	imageGroup := galleryGroup.Group("/image")
	{
		imageGroup.POST("/create", middleware.StaffMiddleware(), galleryController.Create)
		imageGroup.GET("/getall", middleware.OptionalAuthMiddleware(), galleryController.FindAll)
		imageGroup.DELETE("/:uuid", middleware.StaffMiddleware(), galleryController.Delete)
	}
	videoGroup := galleryGroup.Group("/video")
	{
		videoGroup.POST("/create", middleware.StaffMiddleware(), galleryController.CreateVideo)
		videoGroup.GET("/:uuid", galleryController.FindVideoByUUID)
		videoGroup.GET("/getall", middleware.OptionalAuthMiddleware(), galleryController.FindAllVideo)
		videoGroup.DELETE("/:uuid", middleware.StaffMiddleware(), galleryController.DeleteVideo)
	}
}
