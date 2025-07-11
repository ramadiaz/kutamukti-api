package mapper

import (
	"kutamukti-api/api/gallery/dto"
	"kutamukti-api/models"
	"kutamukti-api/pkg/helpers"

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

func MapVideosInputToModel(input dto.Videos) models.Videos {
	var output models.Videos
	mapstructure.Decode(input, &output)
	return output
}

func MapVideosModelToOutput(input models.Videos) dto.VideosResponse {
	var output dto.VideosResponse
	mapstructure.Decode(input, &output)
	youtubeID, _ := helpers.ExtractYouTubeID(input.YoutubeURL)
	output.YoutubeID = youtubeID
	output.ThumbnailURL = helpers.GenerateYoutubeThumbnail(youtubeID)
	return output
}
