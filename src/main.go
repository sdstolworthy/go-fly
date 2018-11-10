package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type SkyscannerResponse struct {
	Quotes []struct {
		QuoteID     int     `json:"QuoteId"`
		MinPrice    float64 `json:"MinPrice"`
		Direct      bool    `json:"Direct"`
		OutboundLeg struct {
			CarrierIds    []int  `json:"CarrierIds"`
			OriginID      int    `json:"OriginId"`
			DestinationID int    `json:"DestinationId"`
			DepartureDate string `json:"DepartureDate"`
		} `json:"OutboundLeg"`
		InboundLeg struct {
			CarrierIds    []int  `json:"CarrierIds"`
			OriginID      int    `json:"OriginId"`
			DestinationID int    `json:"DestinationId"`
			DepartureDate string `json:"DepartureDate"`
		} `json:"InboundLeg"`
		QuoteDateTime string `json:"QuoteDateTime"`
	} `json:"Quotes"`
	Places []struct {
		PlaceID        int    `json:"PlaceId"`
		IataCode       string `json:"IataCode"`
		Name           string `json:"Name"`
		Type           string `json:"Type"`
		SkyscannerCode string `json:"SkyscannerCode"`
		CityName       string `json:"CityName"`
		CityID         string `json:"CityId"`
		CountryName    string `json:"CountryName"`
	} `json:"Places"`
	Carriers []struct {
		CarrierID int    `json:"CarrierId"`
		Name      string `json:"Name"`
	} `json:"Carriers"`
	Currencies []struct {
		Code                        string `json:"Code"`
		Symbol                      string `json:"Symbol"`
		ThousandsSeparator          string `json:"ThousandsSeparator"`
		DecimalSeparator            string `json:"DecimalSeparator"`
		SymbolOnLeft                bool   `json:"SymbolOnLeft"`
		SpaceBetweenAmountAndSymbol bool   `json:"SpaceBetweenAmountAndSymbol"`
		RoundingCoefficient         int    `json:"RoundingCoefficient"`
		DecimalDigits               int    `json:"DecimalDigits"`
	} `json:"Currencies"`
}

func main() {
	url := "https://skyscanner-skyscanner-flight-search-v1.p.mashape.com/apiservices/browsequotes/v1.0/US/USD/en-US/SLC-sky/BNA-sky/2018-12/2019-01"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	flightClient := http.Client{
		Timeout: time.Second * 2,
	}

	req.Header.Set("X-Mashape-Key", "vNqhC5tLOcmshv945nIkMTvjcxh6p1BXD5JjsnZYdq8HsSygJZ")
	req.Header.Set("X-Mashape-Host", "skyscanner-skyscanner-flight-search-v1.p.mashape.com")

	res, getErr := flightClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	apiResponse := SkyscannerResponse{}

	jsonErr := json.Unmarshal(body, &apiResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Printf("%+v\n", apiResponse.Quotes)

}
