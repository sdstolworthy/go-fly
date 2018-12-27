package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sdstolworthy/go-fly/controllers"
	"github.com/sdstolworthy/go-fly/environment"
	skyscanner "github.com/sdstolworthy/go-skyscanner"
)

// Environment contains the application environment

type quoteChannel struct {
	quote *skyscanner.QuoteSummary
	err   error
}

func main() {
	environment.InitializeDatabase()
	defer environment.CloseDatabase()

	router := gin.Default()
	defineRoutes(router)
	router.Run()
}

func defineRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	})
	quoteController := new(controllers.QuoteController)
	quoteController.SetRoutes(router.Group("/quotes"))
}

func test() {

	fmt.Println()

}
