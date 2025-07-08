package services

import (
	"kutamukti-api/api/news/dto"
	"kutamukti-api/api/news/repositories"
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

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.News) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}
	input := mapper.MapNewsInputToModel(data)
	input.UUID = uuid.NewString()
	input.Slug = helpers.FormatSlug(input.Title)
	err := s.repo.Create(ctx, s.DB, input)
	if err != nil {
		return err
	}
	return nil
}

func (s *CompServicesImpl) FindAll(ctx *gin.Context) (*[]dto.NewsResponse, *exceptions.Exception) {
	output, err := s.repo.FindAll(ctx, s.DB)
	if err != nil {
		return nil, err
	}
	var response []dto.NewsResponse
	for _, v := range output {
		response = append(response, mapper.MapNewsModelToOutput(v))
	}
	return &response, nil
}

func (s *CompServicesImpl) FindBySlug(ctx *gin.Context, slug string) (*dto.NewsResponse, *exceptions.Exception) {
	output, err := s.repo.FindBySlug(ctx, s.DB, slug)
	if err != nil {
		return nil, err
	}
	response := mapper.MapNewsModelToOutput(*output)

	return &response, nil
}
func (s *CompServicesImpl) DeleteByUUID(ctx *gin.Context, uuid string) *exceptions.Exception {
	return s.repo.DeleteByUUID(ctx, s.DB, uuid)
}
