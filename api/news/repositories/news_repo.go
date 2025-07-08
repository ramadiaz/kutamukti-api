package repositories

import (
	"kutamukti-api/models"
	"kutamukti-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.News) *exceptions.Exception
	FindAll(ctx *gin.Context, tx *gorm.DB) ([]models.News, *exceptions.Exception)
	FindBySlug(ctx *gin.Context, tx *gorm.DB, slug string) (*models.News, *exceptions.Exception)
	DeleteByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception
}
