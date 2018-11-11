package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/jinzhu/gorm"
	"github.com/sdstolworthy/go-fly/models"
	"github.com/sdstolworthy/go-fly/skyscanner"
)

var db *gorm.DB

// Env contains the application environment
type Env struct {
	db models.Datastore
}

type quoteChannel struct {
	quote skyscanner.QuoteSummary
	err   error
}

func main() {
	db, err := models.NewDB("test.db")
	if err != nil {
		log.Fatal(err)
	}

	env := &Env{db}

	fmt.Println()
	params := skyscanner.Parameters{
		Adults:           1,
		Country:          "US",
		Currency:         "USD",
		Locale:           "en-US",
		OriginPlace:      "SLC-sky",
		DestinationPlace: "BNA-sky",
		OutbandDate:      "anytime",
		InboundDate:      "anytime",
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	quoteChannels := make(chan quoteChannel)

	for _, v := range DestinationAirports {
		go processDestination(v, &params, quoteChannels)
	}
	for range DestinationAirports {
		q := <-quoteChannels
		if err != nil {
			log.Printf("%v\n\n", err)
			continue
		}
		env.db.AddQuote(&models.Quote{
			Price:              q.quote.Price,
			DestinationAirport: params.DestinationPlace,
			OriginAirport:      params.OriginPlace,
		})
		fmt.Fprintf(w, "%v\nPrice:\t$%v\nDeparture:\t%v\nReturn:\t%v\t\n\n", q.quote.DestinationCity, q.quote.Price, q.quote.DepartureDate, q.quote.InboundDate)
	}
}

func processDestination(destination string, params *skyscanner.Parameters, out chan<- quoteChannel) {
	params.DestinationPlace = destination
	SkyscannerQuotes := skyscanner.BrowseQuotes(*params)

	quote, err := SkyscannerQuotes.LowestPrice()
	out <- quoteChannel{
		err:   err,
		quote: *quote,
	}
}
