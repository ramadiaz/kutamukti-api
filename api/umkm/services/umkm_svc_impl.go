package services

import (
	"kutamukti-api/api/umkm/dto"
	"kutamukti-api/api/umkm/repositories"
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

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.UMKM) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	phoneNumber, err := helpers.NormalizePhoneNumber(data.Contact)
	if err != nil {
		return err
	}

	data.Contact = phoneNumber

	input := mapper.MapUMKMInputToModel(data)
	input.UUID = uuid.NewString()
	err = s.repo.Create(ctx, s.DB, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *CompServicesImpl) FindAll(ctx *gin.Context) (*[]dto.UMKMResponse, *exceptions.Exception) {
	output, err := s.repo.FindAll(ctx, s.DB)
	if err != nil {
		return nil, err
	}
	var response []dto.UMKMResponse
	for _, v := range output {
		response = append(response, mapper.MapUMKMModelToOutput(v))
	}
	return &response, nil
}

func (s *CompServicesImpl) FindByUUID(ctx *gin.Context, uuid string) (*dto.UMKMResponse, *exceptions.Exception) {
	output, err := s.repo.FindByUUID(ctx, s.DB, uuid)
	if err != nil {
		return nil, err
	}
	response := mapper.MapUMKMModelToOutput(*output)

	return &response, nil
}

func (s *CompServicesImpl) CreateProduct(ctx *gin.Context, data dto.UMKMProduct) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}
	input := mapper.MapUMKMProductInputToModel(data)
	input.UUID = uuid.NewString()

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	err := s.repo.CreateProduct(ctx, tx, input)
	if err != nil {
		return err
	}
	return nil
}

func (s *CompServicesImpl) FindAllProduct(ctx *gin.Context) (*[]dto.UMKMProductResponse, *exceptions.Exception) {
	output, err := s.repo.FindAllProduct(ctx, s.DB)
	if err != nil {
		return nil, err
	}
	var response []dto.UMKMProductResponse
	for _, v := range output {
		response = append(response, mapper.MapUMKMProductModelToOutput(v))
	}
	return &response, nil
}

func (s *CompServicesImpl) FindProductByKeyword(ctx *gin.Context, keyword string) (*[]dto.UMKMProductResponse, *exceptions.Exception) {
	output, err := s.repo.FindProductByKeyword(ctx, s.DB, keyword)
	if err != nil {
		return nil, err
	}
	var response []dto.UMKMProductResponse
	for _, v := range *output {
		response = append(response, mapper.MapUMKMProductModelToOutput(v))
	}
	return &response, nil
}

func (s *CompServicesImpl) FindProductByUMKMUUID(ctx *gin.Context, uuid string) (*[]dto.UMKMProductResponse, *exceptions.Exception) {
	output, err := s.repo.FindProductByUMKMUUID(ctx, s.DB, uuid)
	if err != nil {
		return nil, err
	}
	var response []dto.UMKMProductResponse
	for _, v := range *output {
		response = append(response, mapper.MapUMKMProductModelToOutput(v))
	}
	return &response, nil
}

func (s *CompServicesImpl) DeleteProduct(ctx *gin.Context, uuid string) *exceptions.Exception {
	return s.repo.DeleteProduct(ctx, s.DB, uuid)
}
