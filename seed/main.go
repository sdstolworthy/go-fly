package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sdstolworthy/go-fly/environment"
	skyscanner "github.com/sdstolworthy/go-skyscanner"
)

func main() {
	airportPtr := flag.Bool("airport", false, "Only seed airports")
	flag.Parse()
	environment.InitializeDatabase()

	airports := seedAirports()
	mashapeKey := os.Getenv("MASHAPE_KEY")
	baseURL := os.Getenv("BASE_URL")
	skyscanner.SetConfig(&skyscanner.Config{
		MashapeKey: &mashapeKey,
		BaseURL:    &baseURL,
	})

	fmt.Println(*airportPtr)
	if *airportPtr != true {
		seedQuotes(airports)
	}
}
