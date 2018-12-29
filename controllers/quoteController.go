package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdstolworthy/go-fly/environment"
	"github.com/sdstolworthy/go-fly/models"
	skyscanner "github.com/sdstolworthy/go-skyscanner"
)

// TODO: Remove
var params = skyscanner.BrowseParameters{
	OriginPlace:      "BNA-sky",
	DestinationPlace: "SLC-sky",
	BaseParameters: skyscanner.BaseParameters{
		Adults:      1,
		Country:     "US",
		Currency:    "USD",
		Locale:      "en-US",
		OutbandDate: "anytime",
		InboundDate: "anytime",
	},
}

type quoteParameters struct {
	skyscanner.BrowseParameters
}

type batchParameters struct {
	skyscanner.BatchBrowseParameters
}

// QuoteController handles all quote routes
type QuoteController struct {
	BaseController
}

// SetRoutes initializes the routes for the Quote Controller
func (c *QuoteController) SetRoutes(router *gin.RouterGroup) {
	c.setRouter(router)

	// GET
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	})
	router.GET("/allQuotes", allQuotes)

	// POST
	router.POST("/getQuote", getQuote)
	router.POST("/batchQuotes", batchQuotes)
}

func getQuote(context *gin.Context) {
	var requestBody quoteParameters
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	quote, err := skyscanner.ProcessDestination(&requestBody.BrowseParameters)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No quote found"})
		return
	}
	environment.Env.Db.AddQuote(&models.Quote{
		Price:              quote.Price,
		DestinationAirport: quote.DestinationCity,
		OriginAirport:      quote.OriginCity,
	})
	context.JSON(http.StatusOK, quote)
	return
}

func batchQuotes(context *gin.Context) {
	var requestBody batchParameters
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	quotes := skyscanner.BatchDestinations(&requestBody.BatchBrowseParameters)
	for _, q := range quotes {
		fmt.Println("test", q)
		environment.Env.Db.AddQuote(&models.Quote{
			Price:              q.Price,
			DestinationAirport: q.DestinationCity,
			OriginAirport:      q.OriginCity,
		})
	}
	context.JSON(http.StatusOK, quotes)
}

func allQuotes(context *gin.Context) {
	quotes, _ := environment.Env.Db.AllQuotes()
	for _, v := range quotes {
		fmt.Println(v)
	}
	context.JSON(http.StatusOK, quotes)
}
