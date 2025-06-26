package mapper

import (
	"kutamukti-api/api/complaint/dto"
	"kutamukti-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapComplaintInputToModel(input dto.Complaint) models.Complaints {
	var output models.Complaints

	mapstructure.Decode(input, &output)
	return output
}
