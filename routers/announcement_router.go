package routers

import (
	"kutamukti-api/api/announcement/controllers"
	"kutamukti-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func AnnouncementRoutes(r *gin.RouterGroup, announcementController controllers.CompControllers) {
	announcementGroup := r.Group("/announcement")
	{
		announcementGroup.POST("/create", middleware.StaffMiddleware(), announcementController.Create)
	}
}
