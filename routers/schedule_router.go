package routers

import (
	"kutamukti-api/api/schedule/controllers"

	"github.com/gin-gonic/gin"
)

func ScheduleRoutes(r *gin.RouterGroup, scheduleController controllers.CompControllers) {
	scheduleGroup := r.Group("/schedule")
	{
		scheduleGroup.POST("/create", scheduleController.Create)
		scheduleGroup.GET("/getall", scheduleController.FindAll)
	}
}
