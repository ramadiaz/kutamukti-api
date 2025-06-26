package repositories

import (
	"kutamukti-api/models"
	"kutamukti-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Schedules) *exceptions.Exception
	FindAll(ctx *gin.Context, tx *gorm.DB) ([]models.Schedules, *exceptions.Exception)
	Update(ctx *gin.Context, tx *gorm.DB, data models.Schedules) *exceptions.Exception
	Delete(ctx *gin.Context, tx *gorm.DB, data models.Schedules) *exceptions.Exception
}
