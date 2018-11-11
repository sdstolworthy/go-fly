package skyscanner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func prettyPrint(apiResponse Response) {
	fmt.Printf("%+v\n", apiResponse.Quotes)
}

func parseResponse(response *http.Response) Response {
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	apiResponse := Response{}
	jsonErr := json.Unmarshal(body, &apiResponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return apiResponse
}

func getRequest(url string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Set("X-Mashape-Key", "vNqhC5tLOcmshv945nIkMTvjcxh6p1BXD5JjsnZYdq8HsSygJZ")
	req.Header.Set("X-Mashape-Host", "skyscanner-skyscanner-flight-search-v1.p.mashape.com")

	if err != nil {
		log.Fatal(err)
	}
	return req
}

func getClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 2,
	}
}

func formatURL(path string) string {
	baseURL := "https://skyscanner-skyscanner-flight-search-v1.p.mashape.com/apiservices/"

	return fmt.Sprintf("%v%v", baseURL, path)
}

func get(url string) *http.Response {
	res, getErr := getClient().Do(getRequest(url))

	if getErr != nil {
		log.Fatal(getErr)
	}

	return res

}

/*
BrowseQuotes stub
*/
func BrowseQuotes(parameters Parameters) Response {
	browseQuotes := formatURL(fmt.Sprintf("browsequotes/v1.0/%v/%v/%v/%v/%v/%v/%v",
		parameters.Country,
		parameters.Currency,
		parameters.Locale,
		parameters.OriginPlace,
		parameters.DestinationPlace,
		parameters.OutbandDate,
		parameters.InboundDate,
	))
	res := get(browseQuotes)
	// fmt.Printf("%+v\n", parsedResponse)
	return parseResponse(res)
}
