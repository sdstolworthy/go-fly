package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/sdstolworthy/go-fly/environment"
	"github.com/sdstolworthy/go-fly/models"
)

func seedAirports() {
	parseAirportCSV()
}

func parseAirportCSV() {
	type Field int
	const (
		id               Field = 0
		ident            Field = 1
		airportType      Field = 2
		name             Field = 3
		latitudeDeg      Field = 4
		longitudeDeg     Field = 5
		elevation        Field = 6
		continent        Field = 7
		isoCountry       Field = 8
		isoRegion        Field = 9
		municipality     Field = 10
		scheduledSerivce Field = 11
		gpsCode          Field = 12
		iataCode         Field = 13
		localCode        Field = 14
		homeLink         Field = 15
		wikipediaLink    Field = 16
		keywords         Field = 17
	)
	// Load a csv file.
	f, _ := os.Open("../data/airports.csv")

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		if record[airportType] != "large_airport" {
			continue
		}
		if record[iataCode] == "BNA" {
			fmt.Println(record[iataCode])
		}
		fmt.Println(record)
		environment.Env.Db.SaveAirport(&models.Airport{
			Continent:    record[continent],
			Country:      record[isoCountry],
			IataCode:     record[iataCode],
			Identifier:   record[ident],
			Region:       record[isoRegion],
			Municipality: record[municipality],
			Name:         record[name],
		})
	}
}
