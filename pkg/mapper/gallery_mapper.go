package mapper

import (
	"kutamukti-api/api/gallery/dto"
	"kutamukti-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapGalleryInputToModel(input dto.Gallery) models.Galleries {
	var output models.Galleries
	mapstructure.Decode(input, &output)
	return output
}

func MapGalleryModelToOutput(input models.Galleries) dto.GalleryResponse {
	var output dto.GalleryResponse
	mapstructure.Decode(input, &output)
	return output
}
