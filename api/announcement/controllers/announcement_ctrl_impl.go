package controllers

import (
	"kutamukti-api/api/announcement/dto"
	"kutamukti-api/api/announcement/services"
	"kutamukti-api/pkg/exceptions"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var announcementFindAllCache struct {
	sync.Mutex
	data *[]dto.AnnouncementResponse
	time time.Time
}

const announcementFindAllCacheDuration = 10 * time.Minute

type CompControllersImpl struct {
	services services.CompServices
}

func NewCompController(compServices services.CompServices) CompControllers {
	return &CompControllersImpl{
		services: compServices,
	}
}

func (h *CompControllersImpl) Create(ctx *gin.Context) {
	var data dto.Announcement

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
	announcementFindAllCache.Lock()
	if announcementFindAllCache.data != nil && time.Since(announcementFindAllCache.time) < announcementFindAllCacheDuration {
		data := announcementFindAllCache.data
		announcementFindAllCache.Unlock()
		ctx.JSON(http.StatusOK, dto.Response{
			Status:  http.StatusOK,
			Message: "success (cached)",
			Body:    data,
		})
		return
	}
	announcementFindAllCache.Unlock()

	data, err := h.services.FindAll(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	announcementFindAllCache.Lock()
	announcementFindAllCache.data = data
	announcementFindAllCache.time = time.Now()
	announcementFindAllCache.Unlock()

	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Body:    data,
	})
}
