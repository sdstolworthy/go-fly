package main

import (
	"fmt"

	"github.com/sdstolworthy/go-scraper/skyscanner"
)

func main() {
	params := skyscanner.Parameters{
		Adults:           1,
		Country:          "US",
		Currency:         "USD",
		Locale:           "en-US",
		OriginPlace:      "SLC-sky",
		DestinationPlace: "BNA-sky",
		OutbandDate:      "2019",
		InboundDate:      "2019",
	}

	for _, v := range DestinationAirports {
		params.DestinationPlace = v
		fmt.Println(v)
		SkyscannerQuotes := skyscanner.BrowseQuotes(params)
		quote := SkyscannerQuotes.LowestPrice()
		fmt.Printf("Price:\t$%v\nDate:\t%v\n", quote.Price, quote.Date)
	}
}
