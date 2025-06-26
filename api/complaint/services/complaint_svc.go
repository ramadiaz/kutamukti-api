package services

import (
	"kutamukti-api/api/complaint/dto"
	"kutamukti-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.Complaint) *exceptions.Exception
	FindAll(ctx *gin.Context) ([]dto.ComplaintResponse, *exceptions.Exception)
}
