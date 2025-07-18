package controllers

import (
	"kutamukti-api/api/users/dto"
	"kutamukti-api/api/users/services"
	"kutamukti-api/pkg/exceptions"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var usersFindAllCache struct {
	sync.Mutex
	data []dto.UserResponse
	time time.Time
}

const usersFindAllCacheDuration = 10 * time.Minute

type CompControllersImpl struct {
	services services.CompServices
}

func NewCompController(compServices services.CompServices) CompControllers {
	return &CompControllersImpl{
		services: compServices,
	}
}

func (h *CompControllersImpl) Create(ctx *gin.Context) {
	var data dto.User

	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.Create(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "success",
	})
}

func (h *CompControllersImpl) SignIn(ctx *gin.Context) {
	var data dto.UserSignIn

	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	token, err := h.services.SignIn(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Body:    token,
	})
}

func (h *CompControllersImpl) FindAll(ctx *gin.Context) {
	data, err := h.services.FindAll(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Body:    data,
	})
}
