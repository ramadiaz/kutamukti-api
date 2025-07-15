package repositories

import (
	"kutamukti-api/models"
	"kutamukti-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.UMKM) *exceptions.Exception
	FindAll(ctx *gin.Context, tx *gorm.DB) ([]models.UMKM, *exceptions.Exception)
	FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.UMKM, *exceptions.Exception)
	Update(ctx *gin.Context, tx *gorm.DB, data models.UMKM) *exceptions.Exception
	Delete(ctx *gin.Context, tx *gorm.DB, data models.UMKM) *exceptions.Exception

	CreateProduct(ctx *gin.Context, tx *gorm.DB, data models.UMKMProduct) *exceptions.Exception
	FindAllProduct(ctx *gin.Context, tx *gorm.DB) ([]models.UMKMProduct, *exceptions.Exception)
	FindProductByKeyword(ctx *gin.Context, tx *gorm.DB, keyword string) (*[]models.UMKMProduct, *exceptions.Exception)
	FindProductByUMKMUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.UMKMProduct, *exceptions.Exception)
	DeleteProduct(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception
}
