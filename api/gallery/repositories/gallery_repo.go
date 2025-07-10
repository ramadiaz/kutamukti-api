package repositories

import (
	"kutamukti-api/models"
	"kutamukti-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.ImageGalleries) *exceptions.Exception
	FindAll(ctx *gin.Context, tx *gorm.DB) ([]models.ImageGalleries, *exceptions.Exception)
	Update(ctx *gin.Context, tx *gorm.DB, data models.ImageGalleries) *exceptions.Exception
	Delete(ctx *gin.Context, tx *gorm.DB, data models.ImageGalleries) *exceptions.Exception
}
