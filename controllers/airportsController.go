package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sdstolworthy/go-fly/environment"
	"github.com/sdstolworthy/go-fly/models"
)

type AirportController struct {
	BaseController
}

func (c *AirportController) SetRoutes(router *gin.RouterGroup) {
	c.setRouter(router)

	// GET
	router.GET("/ping", ping)
	router.GET("/searchAirports", searchAirports)
}

func searchAirports(context *gin.Context) {
	airportCode := context.Query("iataCode")
	if airportCode != "" {
		airportResults, err := environment.Env.Db.SearchAirportsByIATA(&models.Airport{IataCode: airportCode})
		if err != nil {
			context.JSON(http.StatusNoContent, gin.H{"message": "No results found"})
			return
		}
		context.JSON(http.StatusOK, airportResults)
	}
	cityName := context.Query("cityName")
	if cityName != "" {
		airportResults, err := environment.Env.Db.SearchAirportsByCity(&models.Airport{Municipality: cityName})
		if err != nil {
			context.JSON(http.StatusNoContent, gin.H{"message": "No results found"})
			return
		}
		context.JSON(http.StatusOK, airportResults)
	}
}
