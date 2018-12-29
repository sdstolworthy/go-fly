package main

import (
	"fmt"
	"math/rand"

	"github.com/sdstolworthy/go-fly/environment"
	"github.com/sdstolworthy/go-fly/models"
)

func seedQuotes(airports []*models.Airport) {
	for i := 0; i < 100; i++ {
		environment.Env.Db.AddQuote(&models.Quote{
			OriginAirport:      fmt.Sprintf("%s-sky", airports[randomIndex(0, len(airports)-1)].IataCode),
			DestinationAirport: fmt.Sprintf("%s-sky", airports[randomIndex(0, len(airports)-1)].IataCode),
			Price:              float64(randomIndex(100, 500)),
		})
	}
}

func randomIndex(min int, max int) int {
	return rand.Intn(max-min) + min
}
