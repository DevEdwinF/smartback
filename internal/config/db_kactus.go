package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var KDB *gorm.DB

func KactusDB() {
	loadEnvKactus()

	// dsn := buildDSNKactus()
	dsn := "sqlserver://ASISTENCIA:*T3cn0l0g14-*@localhost:5432?database=KACTUS"

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	KDB = db

	log.Println("Conexión establecida con Kactus")
}

func loadEnvKactus() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func buildDSNKactus() string {
	domain := os.Getenv("KACTUS_DB_DOMAIN")
	user := os.Getenv("KACTUS_DB_USER")
	password := os.Getenv("KACTUS_DB_PASSWORD")
	port := os.Getenv("KACTUS_DB_PORT")
	database := os.Getenv("KACTUS_DB_DATABASE")

	dsn := "sqlserver://%s:%s@%s:%s?database=%s"
	return fmt.Sprintf(dsn, user, password, domain, port, database)
}
