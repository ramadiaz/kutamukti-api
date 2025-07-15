package mapper

import (
	"kutamukti-api/api/umkm/dto"
	"kutamukti-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapUMKMInputToModel(input dto.UMKM) models.UMKM {
	var output models.UMKM
	mapstructure.Decode(input, &output)
	output.OpenAt = input.OpenAt
	output.CloseAt = input.CloseAt

	return output
}

func MapUMKMModelToOutput(input models.UMKM) dto.UMKMResponse {
	var output dto.UMKMResponse
	mapstructure.Decode(input, &output)
	output.OpenAt = input.OpenAt
	output.CloseAt = input.CloseAt

	return output
}

func MapUMKMProductInputToModel(input dto.UMKMProduct) models.UMKMProduct {
	var output models.UMKMProduct
	mapstructure.Decode(input, &output)
	return output
}

func MapUMKMProductModelToOutput(input models.UMKMProduct) dto.UMKMProductResponse {
	var output dto.UMKMProductResponse
	mapstructure.Decode(input, &output)
	return output
}
