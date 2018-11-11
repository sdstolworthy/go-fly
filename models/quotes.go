package models

import "github.com/jinzhu/gorm"

// Quote type representing historical entries into database
type Quote struct {
	gorm.Model
	DestinationAirport string
	OriginAirport      string
	Price              float64
}

// AllQuotes gets all quotes
func (db *DB) AllQuotes() ([]*Quote, error) {
	quotes := make([]*Quote, 0)
	db.Find(&quotes)

	return quotes, nil
}

// AddQuote creates a quote
func (db *DB) AddQuote(quote *Quote) (*Quote, error) {
	db.Create(quote)
	return quote, nil
}
