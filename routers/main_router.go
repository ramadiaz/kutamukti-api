package routers

import (
	"net/http"
	"kutamukti-api/injectors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CompRouters(r *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "pong",
		})
	})

	userController := injectors.InitializeUserController(db, validate)
	complaintController := injectors.InitializeComplaintController(db, validate)

	UserRoutes(r, userController)
	ComplaintRoutes(r, complaintController)
}
