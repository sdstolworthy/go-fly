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

func formatURL(url string) string {
	baseURL := "https://skyscanner-skyscanner-flight-search-v1.p.mashape.com/apiservices/"

	return fmt.Sprintf("%vbrowsequotes/v1.0/US/USD/en-US/SLC-sky/BNA-sky/2018-12/2019-01", baseURL)
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
func BrowseQuotes(parameters Parameters) {
	browseQuotes := formatURL("%vbrowsequotes/v1.0/US/USD/en-US/SLC-sky/BNA-sky/2018-12/2019-01")
	fmt.Printf(browseQuotes)
}
