package models

import "time"

// Quote type representing historical entries into database
type Quote struct {
	ID                 uint      `gorm:"primary_key" json:"id"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
	DeletedAt          time.Time `json:"deletedAt"`
	DestinationAirport string    `json:"destinationAirport"`
	OriginAirport      string    `json:"originAirport"`
	Price              float64   `json:"price"`
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

// DeleteQuote deletes a quote from the db
func (db *DB) DeleteQuote(quote *Quote) (bool, error) {
	db.Delete(quote)
	return true, nil
}
