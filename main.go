package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/sdstolworthy/go-fly/controllers"
	"github.com/sdstolworthy/go-fly/environment"
	skyscanner "github.com/sdstolworthy/go-skyscanner"
)

func main() {
	environment.InitializeDatabase()
	defer environment.CloseDatabase()

	mashapeKey := os.Getenv("MASHAPE_KEY")
	baseURL := os.Getenv("BASE_URL")

	skyscanner.SetConfig(&skyscanner.Config{
		MashapeKey: &mashapeKey,
		BaseURL:    &baseURL,
	})
	router := gin.Default()
	router.Use(cors.Default())
	defineRoutes(router)
	router.Run()
}

func defineRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	})
	quoteController := new(controllers.QuoteController)
	quoteController.SetRoutes(router.Group("/quotes"))

	airportController := new(controllers.AirportController)
	airportController.SetRoutes(router.Group("/airports"))

	fulfillmentController := new(controllers.FulfillmentController)
	fulfillmentController.SetRoutes(router.Group("/fulfillment"))
}
