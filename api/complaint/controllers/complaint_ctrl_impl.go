package controllers

import (
	"kutamukti-api/api/complaint/dto"
	"kutamukti-api/api/complaint/services"
	"kutamukti-api/pkg/exceptions"
	"kutamukti-api/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompControllersImpl struct {
	services services.CompServices
}

func NewCompController(compServices services.CompServices) CompControllers {
	return &CompControllersImpl{
		services: compServices,
	}
}

func (h *CompControllersImpl) Create(ctx *gin.Context) {
	var data dto.Complaint

	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	if !helpers.VerifyRecaptcha(data.Captcha) {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Message: "CAPTCHA verification failed",
		})
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

func (h *CompControllersImpl) FindAll(ctx *gin.Context) {
	output, err := h.services.FindAll(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Body:    output,
	})
}

func (h *CompControllersImpl) UpdateStatus(ctx *gin.Context) {
	uuid := ctx.Query("uuid")
	status := ctx.Query("status")

	if uuid == "" || status == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.UpdateStatus(ctx, uuid, dto.ComplaintStatus(status))
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
	})
}
