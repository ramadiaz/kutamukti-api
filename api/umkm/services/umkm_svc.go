package services

import (
	"kutamukti-api/api/umkm/dto"
	"kutamukti-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.UMKM) *exceptions.Exception
	FindAll(ctx *gin.Context) (*[]dto.UMKMResponse, *exceptions.Exception)
	FindByUUID(ctx *gin.Context, uuid string) (*dto.UMKMResponse, *exceptions.Exception)
	CreateProduct(ctx *gin.Context, data dto.UMKMProduct) *exceptions.Exception
	FindAllProduct(ctx *gin.Context) (*[]dto.UMKMProductResponse, *exceptions.Exception)
	FindProductByKeyword(ctx *gin.Context, keyword string) (*[]dto.UMKMProductResponse, *exceptions.Exception)
	FindProductByUMKMUUID(ctx *gin.Context, uuid string) (*[]dto.UMKMProductResponse, *exceptions.Exception)
	DeleteProduct(ctx *gin.Context, uuid string) *exceptions.Exception
}
