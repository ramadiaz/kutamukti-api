package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	GetAnalyticsData(ctx *gin.Context)
}
