package routers

import (
	"kutamukti-api/api/complaint/controllers"

	"github.com/gin-gonic/gin"
)

func ComplaintRoutes(r *gin.RouterGroup, scheduleController controllers.CompControllers) {
	complaintGroup := r.Group("/complaint")
	{
		complaintGroup.POST("/create", scheduleController.Create)
		complaintGroup.GET("/getall", scheduleController.FindAll)
	}
}
