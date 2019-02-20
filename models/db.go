package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// Be ye therefore justified
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Datastore contains all available database actions
type Datastore interface {
	AllQuotes() ([]*Quote, error)
	AddQuote(*Quote) (*Quote, error)
	SaveAirport(*Airport) (*Airport, error)
	SearchAirportsByIATA(*Airport) ([]*Airport, error)
	SearchAirportsByCity(*Airport) ([]*Airport, error)
}

// DB contains a database
type DB struct {
	*gorm.DB
}

// NewDB returns a pointer to a new database connection
func NewDB(dataSourceName string) (*DB, error) {
	dv := getDatabaseEnvironment()
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dv.Host, dv.Port, dv.User, dv.DbName, dv.Password)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Quote{})
	db.AutoMigrate(&Airport{})

	return &DB{db}, nil
}

// DbVariables represents the variables necessary to open a connection to the postgres db
type DbVariables struct {
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
}

func getDatabaseEnvironment() DbVariables {
	fmt.Println("port", os.Getenv("POSTGRES_PORT"))
	return DbVariables{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		DbName:   os.Getenv("DATABASE_NAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}
}
