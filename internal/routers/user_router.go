package routers

import (
	"kutamukti-api/api/users/controllers"
	"kutamukti-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, userController controllers.CompControllers) {
	authGroup := r.Group("/user")
	authGroup.Use(middleware.InternalMiddleware())
	{
		authGroup.POST("/create", userController.Create)
	}
}
