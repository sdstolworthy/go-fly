package main

import (
	"github.com/sdstolworthy/go-fly/environment"
)

func main() {
	environment.InitializeDatabase()
	airports := seedAirports()
	seedQuotes(airports)
}
