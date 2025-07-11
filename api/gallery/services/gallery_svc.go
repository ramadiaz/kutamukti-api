package services

import (
	"kutamukti-api/api/gallery/dto"
	"kutamukti-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.ImageGalleries) *exceptions.Exception
	FindAll(ctx *gin.Context) ([]dto.ImageGalleriesResponse, *exceptions.Exception)
	Delete(ctx *gin.Context, uuid string) *exceptions.Exception
	CreateVideo(ctx *gin.Context, data dto.Videos) *exceptions.Exception
	FindAllVideo(ctx *gin.Context) ([]dto.VideosResponse, *exceptions.Exception)
}
