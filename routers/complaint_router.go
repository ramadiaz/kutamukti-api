package routers

import (
	"kutamukti-api/api/complaint/controllers"

	"github.com/gin-gonic/gin"
)

func ComplaintRoutes(r *gin.RouterGroup, userController controllers.CompControllers) {
	complaintGroup := r.Group("/complaint")
	{
		complaintGroup.POST("/create", userController.Create)
	}
}
