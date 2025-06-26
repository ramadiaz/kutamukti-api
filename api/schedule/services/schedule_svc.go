package services

import (
	"kutamukti-api/api/schedule/dto"
	"kutamukti-api/pkg/exceptions"

	"github.com/gin-gonic/gin"
)

type CompServices interface {
	Create(ctx *gin.Context, data dto.Schedule) *exceptions.Exception
	FindAll(ctx *gin.Context) ([]dto.ScheduleResponse, *exceptions.Exception)
}
