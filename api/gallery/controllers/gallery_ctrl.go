package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Delete(ctx *gin.Context)
	CreateVideo(ctx *gin.Context)
	FindAllVideo(ctx *gin.Context)
	FindVideoByUUID(ctx *gin.Context)
}
