package main

import (
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
		OutbandDate:      "2019-01-01",
	}
	skyscanner.BrowseQuotes(params)
}
