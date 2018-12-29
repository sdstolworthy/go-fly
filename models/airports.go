package models

import "time"

// Airport is an entry recording a database
type Airport struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt"`
	IataCode     string     `json:"iataCode"`
	Identifier   string     `json:"identifier"`
	Country      string     `json:"country"`
	Region       string     `json:"region"`
	Municipality string     `json:"municipality"`
	Continent    string     `json:"continent"`
	Name         string     `json:"name"`
}

// SaveAirport saves an airport to the database
func (db *DB) SaveAirport(airport *Airport) (*Airport, error) {
	db.Where(Airport{Identifier: airport.Identifier}).Assign(airport).FirstOrCreate(&airport)
	return airport, nil
}
