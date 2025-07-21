package controllers

import (
	"kutamukti-api/api/gallery/dto"
	"kutamukti-api/api/gallery/services"
	"kutamukti-api/pkg/exceptions"
	"net/http"
	"sync"
	"time"

	"kutamukti-api/pkg/helpers"

	"github.com/gin-gonic/gin"
)

var galleryFindAllCache struct {
	sync.Mutex
	data *[]dto.ImageGalleriesResponse
	time time.Time
}

const galleryFindAllCacheDuration = 10 * time.Minute

var galleryFindAllVideoCache struct {
	sync.Mutex
	data *[]dto.VideosResponse
	time time.Time
}

const galleryFindAllVideoCacheDuration = 10 * time.Minute

type CompControllersImpl struct {
	services services.CompServices
}

func NewCompController(compServices services.CompServices) CompControllers {
	return &CompControllersImpl{
		services: compServices,
	}
}

func (h *CompControllersImpl) Create(ctx *gin.Context) {
	var data dto.ImageGalleries
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
		data, err := h.services.FindAll(ctx)
		if err != nil {
			ctx.JSON(err.Status, err)
			return
		}
		galleryFindAllCache.Lock()
		galleryFindAllCache.data = &data
		galleryFindAllCache.time = time.Now()
		galleryFindAllCache.Unlock()
		ctx.JSON(http.StatusOK, dto.Response{
			Status:  http.StatusOK,
			Message: "success (recached)",
			Body:    data,
		})
		return
	}
	galleryFindAllCache.Lock()
	if galleryFindAllCache.data != nil && time.Since(galleryFindAllCache.time) < galleryFindAllCacheDuration {
		data := galleryFindAllCache.data
		galleryFindAllCache.Unlock()
		ctx.JSON(http.StatusOK, dto.Response{
			Status:  http.StatusOK,
			Message: "success (cached)",
			Body:    data,
		})
		return
	}
	galleryFindAllCache.Unlock()
	data, err := h.services.FindAll(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	galleryFindAllCache.Lock()
	galleryFindAllCache.data = &data
	galleryFindAllCache.time = time.Now()
	galleryFindAllCache.Unlock()
	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Body:    data,
	})
}

func (h *CompControllersImpl) Delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.Delete(ctx, uuid)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
	})
}

func (h *CompControllersImpl) CreateVideo(ctx *gin.Context) {
	var data dto.Videos
	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.CreateVideo(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "success",
	})
}

func (h *CompControllersImpl) FindAllVideo(ctx *gin.Context) {
	user, _ := helpers.GetUserData(ctx)
	if user.Role == "admin" || user.Role == "staff" {
		data, err := h.services.FindAllVideo(ctx)
		if err != nil {
			ctx.JSON(err.Status, err)
			return
		}
		galleryFindAllVideoCache.Lock()
		galleryFindAllVideoCache.data = &data
		galleryFindAllVideoCache.time = time.Now()
		galleryFindAllVideoCache.Unlock()
		ctx.JSON(http.StatusOK, dto.Response{
			Status:  http.StatusOK,
			Message: "success (recached)",
			Body:    data,
		})
		return
	}
	galleryFindAllVideoCache.Lock()
	if galleryFindAllVideoCache.data != nil && time.Since(galleryFindAllVideoCache.time) < galleryFindAllVideoCacheDuration {
		data := galleryFindAllVideoCache.data
		galleryFindAllVideoCache.Unlock()
		ctx.JSON(http.StatusOK, dto.Response{
			Status:  http.StatusOK,
			Message: "success (cached)",
			Body:    data,
		})
		return
	}
	galleryFindAllVideoCache.Unlock()
	data, err := h.services.FindAllVideo(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	galleryFindAllVideoCache.Lock()
	galleryFindAllVideoCache.data = &data
	galleryFindAllVideoCache.time = time.Now()
	galleryFindAllVideoCache.Unlock()
	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Body:    data,
	})
}

func (h *CompControllersImpl) FindVideoByUUID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	data, err := h.services.FindVideoByUUID(ctx, uuid)
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

func (h *CompControllersImpl) DeleteVideo(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.DeleteVideo(ctx, uuid)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
	})
}
