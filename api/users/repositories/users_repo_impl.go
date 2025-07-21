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

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.Users) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindByUsername(ctx *gin.Context, tx *gorm.DB, username string) (*models.Users, *exceptions.Exception) {
	var output models.Users

	result := tx.Where("username = ?", username).First(&output)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return &output, nil
}

func (r *CompRepositoriesImpl) FindAll(ctx *gin.Context, tx *gorm.DB) ([]models.Users, *exceptions.Exception) {
	var output []models.Users

	result := tx.Order("created_at DESC").Find(&output)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return output, nil
}
