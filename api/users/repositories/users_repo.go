package repositories

import (
	"kutamukti-api/models"
	"kutamukti-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Users) *exceptions.Exception
	FindByUsername(ctx *gin.Context, tx *gorm.DB, username string) (*models.Users, *exceptions.Exception)
	FindAll(ctx *gin.Context, tx *gorm.DB) ([]models.Users, *exceptions.Exception)
}
