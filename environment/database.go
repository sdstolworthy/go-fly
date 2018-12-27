package environment

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/sdstolworthy/go-fly/models"
)

type Environment struct {
	Db models.Datastore
}

var db *gorm.DB

var Env *Environment

func InitializeDatabase() {
	db, err := models.NewDB("test.db")
	if err != nil {
		log.Fatal(err)
	}
	Env = &Environment{db}
}

func CloseDatabase() {
	db.Close()
}
