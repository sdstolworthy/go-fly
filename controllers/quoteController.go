package controllers

import (
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/gin-gonic/gin"
	"github.com/sdstolworthy/go-fly/environment"
	"github.com/sdstolworthy/go-fly/models"
	skyscanner "github.com/sdstolworthy/go-skyscanner"
)

// TODO: Remove
var params = skyscanner.Parameters{
	Adults:           1,
	Country:          "US",
	Currency:         "USD",
	Locale:           "en-US",
	OriginPlace:      "SLC-sky",
	DestinationPlace: "BNA-sky",
	OutbandDate:      "anytime",
	InboundDate:      "anytime",
}

// QuoteController handles all quote routes
type QuoteController struct {
	BaseController
}

func getQuote(context *gin.Context) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	quoteChannels := make(chan *skyscanner.QuoteSummary)
	for _, v := range DestinationAirports {
		go skyscanner.ProcessDestination(v, &params, quoteChannels)
	}
	for range DestinationAirports {
		q := <-quoteChannels
		if q == nil {
			continue
		}
		environment.Env.Db.AddQuote(&models.Quote{
			Price:              q.Price,
			DestinationAirport: q.DestinationCity,
			OriginAirport:      q.OriginCity,
		})
		fmt.Fprintf(w, "%v\nPrice:\t$%v\nDeparture:\t%v\nReturn:\t%v\t\n\n", q.DestinationCity, q.Price, q.DepartureDate, q.InboundDate)
	}
	quotes, _ := environment.Env.Db.AllQuotes()
	// for _, v := range quotes {
	// 	fmt.Printf("City: %v\nPrice: %v\n\n", v.DestinationAirport, v.Price)
	// }
	context.JSON(http.StatusOK, quotes)
}

// SetRoutes initializes the routes for the Quote Controller
func (c *QuoteController) SetRoutes(router *gin.RouterGroup) {
	c.setRouter(router)
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	})
	router.GET("/getQuote", getQuote)
}
