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

	dsn := buildDSNKactus()

	// dsn := "sqlserver://ASISTENCIA:*T3cn0l0g14-*@10.100.0.18:1433?database=KACTUS"

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
	user := os.Getenv("KDB_USER")
	password := os.Getenv("KDB_PASSWORD")
	domain := os.Getenv("KDB_DOMAIN")
	port := os.Getenv("KDB_PORT")
	database := os.Getenv("KDB_DATABASE")

	dsn := "sqlserver://ASISTENCIA:*T3cn0l0g14-*@10.100.0.18:1433?database=KACTUS"
	return fmt.Sprintf(dsn, user, password, domain, port, database)
}
