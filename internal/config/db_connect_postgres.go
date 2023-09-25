package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	loadEnv()

	dsn := buildDSN()

	// dsn := "host=129.213.83.40 user=asistencias password=Smart2023*++ dbname=asistencias port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	DB = db

	log.Println("Conexión establecida con Postgress")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func buildDSN() string {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	return fmt.Sprintf(dsn, dbHost, dbUser, dbPassword, dbName, dbPort)
}
