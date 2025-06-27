package mapper

import (
	"kutamukti-api/api/announcement/dto"
	"kutamukti-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapAnnouncementInputToModel(input dto.Announcement) models.Announcements {
	var output models.Announcements

	mapstructure.Decode(input, &output)
	return output
}

func MapAnnouncementModelToOutput(input models.Announcements) dto.AnnouncementResponse {
	var output dto.AnnouncementResponse

	mapstructure.Decode(input, &output)
	output.CreatedAt = input.CreatedAt
	output.UpdatedAt = input.UpdatedAt

	return output
}
