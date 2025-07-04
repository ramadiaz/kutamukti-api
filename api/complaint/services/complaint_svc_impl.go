package services

import (
	"fmt"
	"kutamukti-api/api/complaint/dto"
	"kutamukti-api/api/complaint/repositories"
	"kutamukti-api/models"
	"kutamukti-api/pkg/exceptions"
	"kutamukti-api/pkg/helpers"
	"kutamukti-api/pkg/logger"
	"kutamukti-api/pkg/mapper"
	"kutamukti-api/pkg/whatsapp"
	"os"
	"time"

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

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.Complaint) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	input := mapper.MapComplaintInputToModel(data)
	input.UUID = uuid.NewString()

	err := s.repo.Create(ctx, tx, input)
	if err != nil {
		return err
	}

	go func() {
		message := fmt.Sprintf(`
*LAPORAN PENGADUAN BARU!*

Dibuat pada: %s
Laporan: %s
Pesan: %s
		`, helpers.FormatIndonesianTime(time.Now()), input.Title, input.Description)

		err := whatsapp.Send(os.Getenv("FONNTE_GROUP_COMPLAINT_ID"), message)
		if err != nil {
			logger.Error("error sending whatsapp: %v", err)
		}
	}()

	return nil
}

func (s *CompServicesImpl) FindAll(ctx *gin.Context) ([]dto.ComplaintResponse, *exceptions.Exception) {
	output, err := s.repo.FindAll(ctx, s.DB)
	if err != nil {
		return nil, err
	}

	var response []dto.ComplaintResponse
	for _, v := range output {
		response = append(response, mapper.MapComplaintModelToOutput(v))
	}

	return response, nil
}

func (s *CompServicesImpl) UpdateStatus(ctx *gin.Context, uuid string, status dto.ComplaintStatus) *exceptions.Exception {
	return s.repo.UpdateStatus(ctx, s.DB, uuid, models.ComplaintStatus(status))
}
