package helpers

import (
	"kutamukti-api/api/users/dto"
	"kutamukti-api/pkg/exceptions"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) (dto.UserResponse, *exceptions.Exception) {
	var result dto.UserResponse
	user_data, _ := c.Get("user")

	result, ok := user_data.(dto.UserResponse)
	if !ok {
		return result, exceptions.NewException(http.StatusUnauthorized, exceptions.ErrInvalidTokenStructure)
	}

	return result, nil
}
