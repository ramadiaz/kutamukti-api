package controllers

import "github.com/gin-gonic/gin"

type CompControllers interface {
	CreateImage(ctx *gin.Context)
}