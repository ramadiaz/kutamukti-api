package services

import (
	"kutamukti-api/api/users/dto"
	"kutamukti-api/api/users/repositories"
	"kutamukti-api/models"
	"kutamukti-api/pkg/exceptions"
	"kutamukti-api/pkg/helpers"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	emailDTO "kutamukti-api/emails/dto"
	emails "kutamukti-api/emails/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CompServicesImpl struct {
	repo     repositories.CompRepositories
	DB       *gorm.DB
	validate *validator.Validate
}

func NewComponentServices(compRepositories repositories.CompRepositories, db *gorm.DB, validate *validator.Validate) CompServices {
	return &CompServicesImpl{
		repo:     compRepositories,
		DB:       db,
		validate: validate,
	}
}

func (s *CompServicesImpl) Create(ctx *gin.Context, data dto.User) *exceptions.Exception {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return exceptions.NewValidationException(validateErr)
	}

	tx := s.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	password, err := helpers.GeneratePassword(12)
	if err != nil {
		return err
	}

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return err
	}

	err = s.repo.Create(ctx, tx, models.Users{
		Email:          data.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		return err
	}

	go func() {
		err := emails.AccountCredentialsEmail(emailDTO.AccountCredentials{
			Name:        data.Name,
			Position:    data.Role.String(),
			Email:       data.Email,
			Username:    data.Username,
			Password:    password,
			LoginURL:    "https://dash.kutamukti.id/auth/login",
			Year:        strconv.Itoa(time.Now().Year()),
			FacebookURL: "https://www.facebook.com/kutamukti.id",
			TwitterURL:  "https://twitter.com/kutamukti_id",
			LinkedinURL: "https://www.linkedin.com/company/kutamukti/",
		})
		if err != nil {
			log.Println(err)
		}
	}()

	return nil
}

func (s *CompServicesImpl) SignIn(ctx *gin.Context, data dto.UserSignIn) (*string, *exceptions.Exception) {
	validateErr := s.validate.Struct(data)
	if validateErr != nil {
		return nil, exceptions.NewValidationException(validateErr)
	}

	user, err := s.repo.FindByUsername(ctx, s.DB, data.Username)
	if err != nil {
		return nil, err
	}

	err = helpers.CheckPasswordHash(data.Password, user.HashedPassword)
	if err != nil {
		return nil, err
	}

	JWT_SECRET := os.Getenv("JWT_SECRET")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["uuid"] = user.UUID
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["username"] = user.Username
	claims["role"] = user.Role

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	secretKey := []byte(JWT_SECRET)
	tokenString, signErr := token.SignedString(secretKey)
	if signErr != nil {
		return nil, exceptions.NewException(http.StatusInternalServerError, exceptions.ErrTokenGenerate)
	}

	return &tokenString, nil
}
