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
	var config Config
	config.getConfig()
	req.Header.Set("X-Mashape-Key", config.MashapeKey)
	req.Header.Set("X-Mashape-Host", config.MashapeHost)

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
	var c Config
	c.getConfig()
	baseURL := c.getConfig().BaseURL

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
	return parseResponse(res)
}
