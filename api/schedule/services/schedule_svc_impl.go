package services

import (
	"kutamukti-api/api/schedule/dto"
	"kutamukti-api/api/schedule/repositories"
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

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.Schedule) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	input := mapper.MapScheduleInputToModel(data)
	input.UUID = uuid.NewString()

	err := s.repo.Create(ctx, tx, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompServicesImpl) FindAll(ctx *gin.Context) ([]dto.ScheduleResponse, *exceptions.Exception) {
	output, err := s.repo.FindAll(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	var response []dto.ScheduleResponse
	for _, v := range output {
		response = append(response, mapper.MapScheduleModelToOutput(v))
	}

	return response, nil
}
