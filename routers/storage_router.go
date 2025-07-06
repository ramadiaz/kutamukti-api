package routers

import (
	"kutamukti-api/api/storages/controllers"

	"github.com/gin-gonic/gin"
)

func StorageRoutes(r *gin.RouterGroup, storageController controllers.CompControllers) {
	storageGroup := r.Group("/storage")
	imageGroup := storageGroup.Group("/image")
	{
		imageGroup.POST("/create", storageController.CreateImage)
	}
}
