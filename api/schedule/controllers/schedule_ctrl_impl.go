package controllers

import (
	"kutamukti-api/api/schedule/dto"
	"kutamukti-api/api/schedule/services"
	"kutamukti-api/pkg/exceptions"
	"net/http"
	"sync"
	"time"

	"kutamukti-api/pkg/helpers"

	"github.com/gin-gonic/gin"
)

var scheduleFindAllCache struct {
	sync.Mutex
	data []dto.ScheduleResponse
	time time.Time
}

const scheduleFindAllCacheDuration = 10 * time.Minute

type CompControllersImpl struct {
	services services.CompServices
}

func NewCompController(compServices services.CompServices) CompControllers {
	return &CompControllersImpl{
		services: compServices,
	}
}

func (h *CompControllersImpl) Create(ctx *gin.Context) {
	var data dto.Schedule
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
	user, _ := helpers.GetUserData(ctx)
	if user.Role == "admin" || user.Role == "staff" {
		output, err := h.services.FindAll(ctx)
		if err != nil {
			ctx.JSON(err.Status, err)
			return
		}
		scheduleFindAllCache.Lock()
		scheduleFindAllCache.data = output
		scheduleFindAllCache.time = time.Now()
		scheduleFindAllCache.Unlock()
		ctx.JSON(http.StatusOK, dto.Response{
			Status:  http.StatusOK,
			Message: "success (recached)",
			Body:    output,
		})
		return
	}
	scheduleFindAllCache.Lock()
	if scheduleFindAllCache.data != nil && time.Since(scheduleFindAllCache.time) < scheduleFindAllCacheDuration {
		data := scheduleFindAllCache.data
		scheduleFindAllCache.Unlock()
		ctx.JSON(http.StatusOK, dto.Response{
			Status:  http.StatusOK,
			Message: "success (cached)",
			Body:    data,
		})
		return
	}
	scheduleFindAllCache.Unlock()
	output, err := h.services.FindAll(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	scheduleFindAllCache.Lock()
	scheduleFindAllCache.data = output
	scheduleFindAllCache.time = time.Now()
	scheduleFindAllCache.Unlock()
	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Body:    output,
	})
}
