package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// FulfillmentController handles fulfillments from external services
type FulfillmentController struct {
	BaseController
}

// SetRoutes handles setting context for the Fulfillment Controller
func (c *FulfillmentController) SetRoutes(router *gin.RouterGroup) {
	c.setRouter(router)

	// GET
	router.GET("/ping", ping)

	// POST
	router.POST("/fulfillment/dialogflow", dialogflow)
}

func dialogflow(context *gin.Context) {
	fmt.Println("fulfillment")
}
