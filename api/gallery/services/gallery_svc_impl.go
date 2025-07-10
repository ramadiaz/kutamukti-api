package services

import (
	"kutamukti-api/api/gallery/dto"
	"kutamukti-api/api/gallery/repositories"
	"kutamukti-api/pkg/exceptions"
	"kutamukti-api/pkg/helpers"
	"kutamukti-api/pkg/mapper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.ImageGalleries) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	input := mapper.MapImageGalleriesInputToModel(data)
	input.UUID = uuid.NewString()

	err := s.repo.Create(ctx, tx, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompServicesImpl) FindAll(ctx *gin.Context) ([]dto.ImageGalleriesResponse, *exceptions.Exception) {
	output, err := s.repo.FindAll(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	var response []dto.ImageGalleriesResponse
	for _, v := range output {
		result := mapper.MapImageGalleriesModelToOutput(v)
		response = append(response, result)
	}

	return response, nil
}
