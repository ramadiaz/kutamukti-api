package routers

import (
	mainInjectors "kutamukti-api/injectors"
	"kutamukti-api/internal/injectors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func InternalRouters(r *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	internalController := injectors.InitializeAuthController(validate)
	userController := mainInjectors.InitializeUserController(db, validate)

	AuthRoutes(r, internalController)
	UserRoutes(r, userController)
}
