package environment

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/sdstolworthy/go-fly/models"
)

// Environment defines global environment variables
type Environment struct {
	Db models.Datastore
}

var db *gorm.DB

// Env contains global environment variables
var Env *Environment

// InitializeDatabase creates or opens the database
func InitializeDatabase() {
	db, err := models.NewDB("../test.db")
	if err != nil {
		log.Fatal(err)
	}
	Env = &Environment{db}
}

// CloseDatabase closes the database
func CloseDatabase() {
	db.Close()
}
