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

	for _, v := range DestinationAirports {
		params.DestinationPlace = v
		fmt.Println(v)
		SkyscannerQuotes := skyscanner.BrowseQuotes(params)

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		quote, err := SkyscannerQuotes.LowestPrice()
		if err != nil {
			log.Printf("%v\n\n", err)
			continue
		}
		env.db.AddQuote(&models.Quote{
			Price:              quote.Price,
			DestinationAirport: params.DestinationPlace,
			OriginAirport:      params.OriginPlace,
		})

		fmt.Fprintf(w, "Price:\t$%v\nDeparture:\t%v\nReturn:\t%v\t\n\n", quote.Price, quote.DepartureDate, quote.InboundDate)
	}
}
