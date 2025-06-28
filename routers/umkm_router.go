package routers

import (
	"kutamukti-api/api/umkm/controllers"

	"github.com/gin-gonic/gin"
)

func UMKMRoutes(r *gin.RouterGroup, umkmController controllers.CompControllers) {
	umkmGroup := r.Group("/umkm")
	{
		umkmGroup.POST("/create", umkmController.Create)
	}
}
