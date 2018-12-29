package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdstolworthy/go-fly/environment"
	"github.com/sdstolworthy/go-fly/models"
	skyscanner "github.com/sdstolworthy/go-skyscanner"
)

// QuoteController handles all quote routes
type QuoteController struct {
	BaseController
}

// SetRoutes initializes the routes for the Quote Controller
func (c *QuoteController) SetRoutes(router *gin.RouterGroup) {
	c.setRouter(router)

	// GET
	router.GET("/ping", ping)
	router.GET("/allQuotes", allQuotes)

	// POST
	router.POST("/getQuote", getQuote)
	router.POST("/batchQuotes", batchQuotes)
}

type quoteParameters struct {
	skyscanner.BrowseParameters
}

func getQuote(context *gin.Context) {
	var requestBody quoteParameters
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quote, err := skyscanner.ProcessDestination(&requestBody.BrowseParameters)

	if err != nil {
		context.JSON(http.StatusNoContent, gin.H{"message": "No quote found"})
		return
	}
	saveQuote(quote)
	context.JSON(http.StatusOK, quote)
	return
}

type batchParameters struct {
	skyscanner.BatchBrowseParameters
}

func batchQuotes(context *gin.Context) {
	var requestBody batchParameters
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requestBody.BatchBrowseParameters.Destinations = DestinationAirports[:]
	quotes := skyscanner.BatchDestinations(&requestBody.BatchBrowseParameters)
	saveQuotes(quotes)
	context.JSON(http.StatusOK, quotes)
}

func allQuotes(context *gin.Context) {
	quotes, _ := environment.Env.Db.AllQuotes()
	context.JSON(http.StatusOK, quotes)
}

func saveQuote(q *skyscanner.QuoteSummary) {
	environment.Env.Db.AddQuote(&models.Quote{
		Price:              q.Price,
		DestinationAirport: q.DestinationCity,
		OriginAirport:      q.OriginCity,
	})
}

func saveQuotes(q []*skyscanner.QuoteSummary) {
	for _, quote := range q {
		saveQuote(quote)
	}
}
