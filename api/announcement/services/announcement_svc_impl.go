package services

import (
	"fmt"
	"kutamukti-api/api/announcement/dto"
	"kutamukti-api/api/announcement/repositories"
	"kutamukti-api/pkg/exceptions"
	"kutamukti-api/pkg/helpers"
	"kutamukti-api/pkg/logger"
	"kutamukti-api/pkg/mapper"
	"kutamukti-api/pkg/whatsapp"
	"os"

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

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.Announcement) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	input := mapper.MapAnnouncementInputToModel(data)
	input.UUID = uuid.NewString()
	input.Slug = helpers.FormatSlug(input.Title)

	err := s.repo.Create(ctx, s.DB, input)
	if err != nil {
		return err
	}

	go func() {
		message := fmt.Sprintf(`
*PENGUMUMAN BARU!*

%s

%s
		`, input.Title, input.Description)

		err := whatsapp.Send(os.Getenv("FONNTE_GROUP_ANNOUNCEMENT_ID"), message)
		if err != nil {
			logger.Error("error sending whatsapp: %v", err)
		}
	}()

	return nil
}
