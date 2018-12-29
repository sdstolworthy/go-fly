package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/sdstolworthy/go-fly/controllers"
	"github.com/sdstolworthy/go-fly/environment"
)

func main() {
	environment.InitializeDatabase()
	defer environment.CloseDatabase()

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
}
