package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	connStr := "host=localhost user=postgres password=password dbname=go-postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection established")

	return db
}
