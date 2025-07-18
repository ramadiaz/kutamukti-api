package routers

import (
	"kutamukti-api/api/schedule/controllers"
	"kutamukti-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func ScheduleRoutes(r *gin.RouterGroup, scheduleController controllers.CompControllers) {
	scheduleGroup := r.Group("/schedule")
	{
		scheduleGroup.POST("/create", middleware.StaffMiddleware(), scheduleController.Create)
		scheduleGroup.GET("/getall", middleware.OptionalAuthMiddleware(), scheduleController.FindAll)
	}
}
