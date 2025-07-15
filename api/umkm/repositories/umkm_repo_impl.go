package repositories

import (
	"kutamukti-api/models"
	"kutamukti-api/pkg/exceptions"

	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositoriesImpl struct {
}

func NewComponentRepository() CompRepositories {
	return &CompRepositoriesImpl{}
}

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.UMKM) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindAll(ctx *gin.Context, tx *gorm.DB) ([]models.UMKM, *exceptions.Exception) {
	var output []models.UMKM

	result := tx.Find(&output)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return output, nil
}

func (r *CompRepositoriesImpl) FindByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.UMKM, *exceptions.Exception) {
	var output models.UMKM

	result := tx.Preload("Products").Preload("Products.Images").Where("uuid = ?", uuid).First(&output)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &output, nil
}

func (r *CompRepositoriesImpl) Update(ctx *gin.Context, tx *gorm.DB, data models.UMKM) *exceptions.Exception {
	result := tx.Save(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) Delete(ctx *gin.Context, tx *gorm.DB, data models.UMKM) *exceptions.Exception {
	result := tx.Delete(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}
func (r *CompRepositoriesImpl) CreateProduct(ctx *gin.Context, tx *gorm.DB, data models.UMKMProduct) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindAllProduct(ctx *gin.Context, tx *gorm.DB) ([]models.UMKMProduct, *exceptions.Exception) {
	var output []models.UMKMProduct
	result := tx.Preload("Images").Find(&output)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}
	return output, nil
}

func (r *CompRepositoriesImpl) FindProductByKeyword(ctx *gin.Context, tx *gorm.DB, keyword string) (*[]models.UMKMProduct, *exceptions.Exception) {
	var output []models.UMKMProduct
	lowerKeyword := "%" + strings.ToLower(keyword) + "%"
	query := tx.Preload("Images").Preload("UMKM").Where(
		"LOWER(name) LIKE ? OR LOWER(description) LIKE ? OR LOWER(variation) LIKE ? OR CAST(price AS CHAR) LIKE ?",
		lowerKeyword, lowerKeyword, lowerKeyword, lowerKeyword,
	)
	result := query.Find(&output)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}
	return &output, nil
}

func (r *CompRepositoriesImpl) FindProductByUMKMUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*[]models.UMKMProduct, *exceptions.Exception) {
	var output []models.UMKMProduct
	result := tx.Where("umkm_uuid = ?", uuid).Find(&output)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &output, nil
}

func (r *CompRepositoriesImpl) DeleteProduct(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception {
	result := tx.Where("uuid = ?", uuid).Delete(&models.UMKMProduct{})
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}
