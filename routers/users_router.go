package routers

import (
	"kutamukti-api/api/users/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, userController controllers.CompControllers) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/create", userController.Create)
	}
}
