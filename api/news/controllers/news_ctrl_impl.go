package controllers

import (
	"kutamukti-api/api/news/dto"
	"kutamukti-api/api/news/services"
	"kutamukti-api/pkg/exceptions"
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
	var data dto.News

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

func (h *CompControllersImpl) FindBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	data, err := h.services.FindBySlug(ctx, slug)
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

func (h *CompControllersImpl) DeleteByUUID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.DeleteByUUID(ctx, uuid)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
	})
}
