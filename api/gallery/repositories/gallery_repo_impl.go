package repositories

import (
	"kutamukti-api/models"
	"kutamukti-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositoriesImpl struct {
}

func NewComponentRepository() CompRepositories {
	return &CompRepositoriesImpl{}
}

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.ImageGalleries) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindAll(ctx *gin.Context, tx *gorm.DB) ([]models.ImageGalleries, *exceptions.Exception) {
	var output []models.ImageGalleries

	result := tx.
		Preload("Images").
		Find(&output).
		Order("created_at DESC")
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return output, nil
}

func (r *CompRepositoriesImpl) Update(ctx *gin.Context, tx *gorm.DB, data models.ImageGalleries) *exceptions.Exception {
	result := tx.Save(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) Delete(ctx *gin.Context, tx *gorm.DB, uuid string) *exceptions.Exception {
	result := tx.Where("uuid = ?", uuid).Select("Images").Delete(&models.ImageGalleries{})
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) CreateVideo(ctx *gin.Context, tx *gorm.DB, data models.Videos) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindAllVideo(ctx *gin.Context, tx *gorm.DB) ([]models.Videos, *exceptions.Exception) {
	var output []models.Videos

	result := tx.Find(&output)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return output, nil
}

func (r *CompRepositoriesImpl) FindVideoByUUID(ctx *gin.Context, tx *gorm.DB, uuid string) (*models.Videos, *exceptions.Exception) {
	var output models.Videos

	result := tx.Where("uuid = ?", uuid).First(&output)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &output, nil
}
