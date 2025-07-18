package controllers

import (
	"kutamukti-api/api/umkm/dto"
	"kutamukti-api/api/umkm/services"
	"kutamukti-api/pkg/exceptions"
	"net/http"
	"sync"
	"time"

	"kutamukti-api/pkg/helpers"

	"github.com/gin-gonic/gin"
)

var umkmFindAllCache struct {
	sync.Mutex
	data *[]dto.UMKMResponse
	time time.Time
}

const umkmFindAllCacheDuration = 10 * time.Minute

var umkmFindAllProductCache struct {
	sync.Mutex
	data *[]dto.UMKMProductResponse
	time time.Time
}

const umkmFindAllProductCacheDuration = 10 * time.Minute

type CompControllersImpl struct {
	services services.CompServices
}

func NewCompController(compServices services.CompServices) CompControllers {
	return &CompControllersImpl{
		services: compServices,
	}
}

func (h *CompControllersImpl) Create(ctx *gin.Context) {
	var data dto.UMKM
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
		umkmFindAllCache.Lock()
		umkmFindAllCache.data = data
		umkmFindAllCache.time = time.Now()
		umkmFindAllCache.Unlock()
		ctx.JSON(http.StatusOK, dto.Response{
			Status:  http.StatusOK,
			Message: "success (recached)",
			Body:    data,
		})
		return
	}
	umkmFindAllCache.Lock()
	if umkmFindAllCache.data != nil && time.Since(umkmFindAllCache.time) < umkmFindAllCacheDuration {
		data := umkmFindAllCache.data
		umkmFindAllCache.Unlock()
		ctx.JSON(http.StatusOK, dto.Response{
			Status:  http.StatusOK,
			Message: "success (cached)",
			Body:    data,
		})
		return
	}
	umkmFindAllCache.Unlock()
	data, err := h.services.FindAll(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	umkmFindAllCache.Lock()
	umkmFindAllCache.data = data
	umkmFindAllCache.time = time.Now()
	umkmFindAllCache.Unlock()
	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Body:    data,
	})
}

func (h *CompControllersImpl) FindByUUID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	data, err := h.services.FindByUUID(ctx, uuid)
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

func (h *CompControllersImpl) CreateProduct(ctx *gin.Context) {
	var data dto.UMKMProduct
	jsonErr := ctx.ShouldBindJSON(&data)
	if jsonErr != nil {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}

	err := h.services.CreateProduct(ctx, data)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "success",
	})
}

func (h *CompControllersImpl) FindAllProduct(ctx *gin.Context) {
	user, _ := helpers.GetUserData(ctx)
	if user.Role == "admin" || user.Role == "staff" {
		data, err := h.services.FindAllProduct(ctx)
		if err != nil {
			ctx.JSON(err.Status, err)
			return
		}
		umkmFindAllProductCache.Lock()
		umkmFindAllProductCache.data = data
		umkmFindAllProductCache.time = time.Now()
		umkmFindAllProductCache.Unlock()
		ctx.JSON(http.StatusOK, dto.Response{
			Status:  http.StatusOK,
			Message: "success (recached)",
			Body:    data,
		})
		return
	}
	umkmFindAllProductCache.Lock()
	if umkmFindAllProductCache.data != nil && time.Since(umkmFindAllProductCache.time) < umkmFindAllProductCacheDuration {
		data := umkmFindAllProductCache.data
		umkmFindAllProductCache.Unlock()
		ctx.JSON(http.StatusOK, dto.Response{
			Status:  http.StatusOK,
			Message: "success (cached)",
			Body:    data,
		})
		return
	}
	umkmFindAllProductCache.Unlock()
	data, err := h.services.FindAllProduct(ctx)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	umkmFindAllProductCache.Lock()
	umkmFindAllProductCache.data = data
	umkmFindAllProductCache.time = time.Now()
	umkmFindAllProductCache.Unlock()
	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
		Body:    data,
	})
}

func (h *CompControllersImpl) FindProductByKeyword(ctx *gin.Context) {
	keyword := ctx.Query("keyword")

	data, err := h.services.FindProductByKeyword(ctx, keyword)
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

func (h *CompControllersImpl) FindProductByUMKMUUID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}
	data, err := h.services.FindProductByUMKMUUID(ctx, uuid)
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

func (h *CompControllersImpl) DeleteProduct(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, exceptions.NewException(http.StatusBadRequest, exceptions.ErrBadRequest))
		return
	}
	err := h.services.DeleteProduct(ctx, uuid)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "success",
	})
}
