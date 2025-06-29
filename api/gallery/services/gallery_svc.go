package services

import (
	"kutamukti-api/api/gallery/dto"
	"kutamukti-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.Gallery) *exceptions.Exception
}