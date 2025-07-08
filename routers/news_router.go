package routers

import (
	"kutamukti-api/api/news/controllers"
	"kutamukti-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func NewsRoutes(r *gin.RouterGroup, newsController controllers.CompControllers) {
	newsGroup := r.Group("/news")
	{
		newsGroup.POST("/create", middleware.StaffMiddleware(), newsController.Create)
		newsGroup.GET("/getall", newsController.FindAll)
		newsGroup.GET("/:slug", newsController.FindBySlug)
		newsGroup.DELETE("/:uuid", middleware.StaffMiddleware(), newsController.DeleteByUUID)
	}
}
