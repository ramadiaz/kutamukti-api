package routers

import (
	"kutamukti-api/api/gallery/controllers"
	"kutamukti-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func GalleryRoutes(r *gin.RouterGroup, galleryController controllers.CompControllers) {
	galleryGroup := r.Group("/gallery")
	{
		imageGroup := galleryGroup.Group("/image")
		imageGroup.POST("/create", middleware.StaffMiddleware(), galleryController.Create)
		imageGroup.GET("/getall", galleryController.FindAll)
		imageGroup.DELETE("/:uuid", galleryController.Delete)
	}
}
