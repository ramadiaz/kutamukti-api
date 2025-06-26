package repositories

import (
	"kutamukti-api/models"
	"kutamukti-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Complaints) *exceptions.Exception
	FindAll(ctx *gin.Context, tx *gorm.DB) ([]models.Complaints, *exceptions.Exception)
}
