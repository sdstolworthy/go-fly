package models

import (
	"github.com/jinzhu/gorm"
	// Be ye therefore justified
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Datastore contains all available database actions
type Datastore interface {
	AllQuotes() ([]*Quote, error)
	AddQuote(*Quote) (*Quote, error)
	SaveAirport(*Airport) (*Airport, error)
}

// DB contains a database
type DB struct {
	*gorm.DB
}

// NewDB returns a pointer to a new database connection
func NewDB(dataSourceName string) (*DB, error) {
	db, err := gorm.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Quote{})
	db.AutoMigrate(&Airport{})

	return &DB{db}, nil
}
