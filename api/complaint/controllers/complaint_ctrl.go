package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	UpdateStatus(ctx *gin.Context)
}
