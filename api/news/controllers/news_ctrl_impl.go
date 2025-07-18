package controllers

import (
	"kutamukti-api/api/news/dto"
	"kutamukti-api/api/news/services"
	"kutamukti-api/pkg/exceptions"
	"net/http"
	"sync"
	"time"

	"kutamukti-api/pkg/helpers"

	"github.com/gin-gonic/gin"
)

var newsFindAllCache struct {
	sync.Mutex
	data *[]dto.NewsResponse
	time time.Time
}

const newsFindAllCacheDuration = 10 * time.Minute

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

	result, err := h.services.Create(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Body:    result,
		Message: "success",
	})
}

func (h *CompControllersImpl) FindAll(ctx *gin.Context) {
	user, _ := helpers.GetUserData(ctx)
	if user.Role == "admin" || user.Role == "staff" {
		data, err := h.services.FindAll(ctx)
		if err != nil {
			ctx.JSON(err.Status, err)
			return
		}
		newsFindAllCache.Lock()
		newsFindAllCache.data = data
		newsFindAllCache.time = time.Now()
		newsFindAllCache.Unlock()
		ctx.JSON(http.StatusOK, dto.Response{
			Status:  http.StatusOK,
			Message: "success (recached)",
			Body:    data,
		})
		return
	}
	newsFindAllCache.Lock()
	if newsFindAllCache.data != nil && time.Since(newsFindAllCache.time) < newsFindAllCacheDuration {
		data := newsFindAllCache.data
		newsFindAllCache.Unlock()
		ctx.JSON(http.StatusOK, dto.Response{
			Status:  http.StatusOK,
			Message: "success (cached)",
			Body:    data,
		})
		return
	}
	newsFindAllCache.Unlock()
	data, err := h.services.FindAll(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	newsFindAllCache.Lock()
	newsFindAllCache.data = data
	newsFindAllCache.time = time.Now()
	newsFindAllCache.Unlock()
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
