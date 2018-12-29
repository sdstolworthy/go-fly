package controllers

import (
	"net/http"

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

func ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
