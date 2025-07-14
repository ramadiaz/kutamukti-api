package routers

import (
	"kutamukti-api/api/umkm/controllers"
	"kutamukti-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func UMKMRoutes(r *gin.RouterGroup, umkmController controllers.CompControllers) {
	umkmGroup := r.Group("/umkm")
	{
		umkmGroup.POST("/create", middleware.StaffMiddleware(), umkmController.Create)
		umkmGroup.GET("/getall", umkmController.FindAll)
		umkmGroup.GET("/:uuid", umkmController.FindByUUID)
	}
}
