package services

import (
	"kutamukti-api/api/news/dto"
	"kutamukti-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.News) *exceptions.Exception
}
