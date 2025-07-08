package mapper

import (
	"kutamukti-api/api/news/dto"
	"kutamukti-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapNewsInputToModel(input dto.News) models.News {
	var output models.News
	mapstructure.Decode(input, &output)
	return output
}

func MapNewsModelToOutput(input models.News) dto.NewsResponse {
	var output dto.NewsResponse
	mapstructure.Decode(input, &output)
	output.CreatedAt = input.CreatedAt
	return output
}
