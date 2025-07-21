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

func (r *CompRepositoriesImpl) Create(ctx *gin.Context, tx *gorm.DB, data models.Complaints) *exceptions.Exception {
	result := tx.Create(&data)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}

func (r *CompRepositoriesImpl) FindAll(ctx *gin.Context, tx *gorm.DB) ([]models.Complaints, *exceptions.Exception) {
	var output []models.Complaints

	result := tx.Order("created_at DESC").Find(&output)
	if result.Error != nil {
		return nil, exceptions.ParseGormError(tx, result.Error)
	}

	return output, nil
}

func (r *CompRepositoriesImpl) UpdateStatus(ctx *gin.Context, tx *gorm.DB, uuid string, status models.ComplaintStatus) *exceptions.Exception {
	result := tx.Model(&models.Complaints{}).Where("uuid = ?", uuid).Update("status", status)
	if result.Error != nil {
		return exceptions.ParseGormError(tx, result.Error)
	}

	return nil
}
