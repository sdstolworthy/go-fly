package main

import (
	"flag"
	"fmt"

	"github.com/sdstolworthy/go-fly/environment"
)

func main() {
	airportPtr := flag.Bool("airport", false, "Only seed airports")
	environment.InitializeDatabase()

	airports := seedAirports()

	flag.Parse()

	fmt.Println(*airportPtr)
	if *airportPtr != true {
		seedQuotes(airports)
	}
}
