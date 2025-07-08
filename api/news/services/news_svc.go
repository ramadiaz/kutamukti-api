package services

import (
	"kutamukti-api/api/news/dto"
	"kutamukti-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.News) *exceptions.Exception
	FindAll(ctx *gin.Context) (*[]dto.NewsResponse, *exceptions.Exception)
	FindBySlug(ctx *gin.Context, slug string) (*dto.NewsResponse, *exceptions.Exception)
	DeleteByUUID(ctx *gin.Context, uuid string) *exceptions.Exception
}
