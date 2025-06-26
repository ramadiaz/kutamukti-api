package services

import (
	"kutamukti-api/api/users/dto"
	"kutamukti-api/api/users/repositories"
	"kutamukti-api/models"
	"kutamukti-api/pkg/exceptions"
	"kutamukti-api/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	validate *validator.Validate
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate) CompServices {
	return &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		validate: validate,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.Users) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	hashedPassword, err := helpers.HashPassword(data.Passoword)
	if err != nil {
		return err
	}

	err = s.repo.Create(ctx, tx, models.Users{
		Email:          data.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		return err
	}

	return nil
}
