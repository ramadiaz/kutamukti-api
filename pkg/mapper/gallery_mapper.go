package mapper

import (
	"kutamukti-api/api/gallery/dto"
	"kutamukti-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapImageGalleriesInputToModel(input dto.ImageGalleries) models.ImageGalleries {
	var output models.ImageGalleries
	mapstructure.Decode(input, &output)
	return output
}

func MapImageGalleriesModelToOutput(input models.ImageGalleries) dto.ImageGalleriesResponse {
	var output dto.ImageGalleriesResponse
	mapstructure.Decode(input, &output)
	output.CreatedAt = input.CreatedAt
	return output
}
