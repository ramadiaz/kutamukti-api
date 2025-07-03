package routers

import (
	"kutamukti-api/api/users/controllers"
	"kutamukti-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, userController controllers.CompControllers) {
	userGroup := r.Group("/user")
	{
		userGroup.POST("/create", middleware.AdminMiddleware(), userController.Create)
		userGroup.POST("/signin", userController.SignIn)
	}
}
