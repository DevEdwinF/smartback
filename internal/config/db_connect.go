package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=containers-us-west-210.railway.app user=postgres password=7xt3Vx6eAevhZTMmSiGJ dbname=railway port=7112 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Database connection successfully established")
}
