package mapper

import (
	"kutamukti-api/api/users/dto"
	"kutamukti-api/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapUserInputToModel(input dto.User) models.Users {
	var user models.Users

	mapstructure.Decode(input, &user)
	return user
}
