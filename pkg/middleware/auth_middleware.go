package middleware

import (
	"kutamukti-api/api/users/dto"
	"kutamukti-api/pkg/exceptions"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func StaffMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := os.Getenv("JWT_SECRET")

		var secretKey = []byte(secret)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrForbidden))
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidCredentials))
			return
		}

		tokenString := authHeaderParts[1]
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidCredentials))
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidCredentials))
			return
		}

		user := dto.UserResponse{
			UUID:     claims["uuid"].(string),
			Email:    claims["email"].(string),
			Name:     claims["name"].(string),
			Username: claims["username"].(string),
			Role:     dto.Roles(claims["role"].(string)),
		}

		c.Set("user", user)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := os.Getenv("JWT_SECRET")

		var secretKey = []byte(secret)

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrForbidden))
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidCredentials))
			return
		}

		tokenString := authHeaderParts[1]
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidCredentials))
			return
		}

		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrInvalidCredentials))
			return
		}

		user := dto.UserResponse{
			UUID:     claims["uuid"].(string),
			Email:    claims["email"].(string),
			Name:     claims["name"].(string),
			Username: claims["username"].(string),
			Role:     dto.Roles(claims["role"].(string)),
		}

		if user.Role != dto.Admin {
			c.AbortWithStatusJSON(http.StatusForbidden, exceptions.NewException(http.StatusForbidden, exceptions.ErrForbidden))
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
