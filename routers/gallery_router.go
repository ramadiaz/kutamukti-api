package routers

import (
	"kutamukti-api/api/gallery/controllers"

	"github.com/gin-gonic/gin"
)

func GalleryRoutes(r *gin.RouterGroup, galleryController controllers.CompControllers) {
	galleryGroup := r.Group("/gallery")
	{
		galleryGroup.POST("/create", galleryController.Create)
	}
}
