package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes the SQLite database connection
func Connect() {
	db, err := gorm.Open(sqlite.Open("../payments.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	DB = db
	log.Println("SQLite Database connected!")
}
