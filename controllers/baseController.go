package controllers

import (
	"github.com/gin-gonic/gin"
)

type baseController interface {
	setRouter(router *gin.RouterGroup)
}

type BaseController struct {
	baseController
	router *gin.RouterGroup
}

func (c *BaseController) setRouter(router *gin.RouterGroup) {
	c.router = router
}
