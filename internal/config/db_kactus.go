// package config

// import (
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/sqlserver"
// 	"gorm.io/gorm"
// )

// var KDB *gorm.DB

// func KactusDB() {
// 	loadEnvKactus()

// 	dsn := KactusDSN()

// 	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Error connecting to database: %v", err)
// 	}

// 	KDB = db

// 	log.Println("Database connection successfully established")
// }

// func loadEnvKactus() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}
// }

// func KactusDSN() string {
// 	domain := os.Getenv("KDB_DOMAIN")
// 	user := os.Getenv("KDB_USER")
// 	pass := os.Getenv("KDB_PASS")
// 	port := os.Getenv("KDB_PORT")
// 	database := os.Getenv("KDB_DATABASE")

// 	dsn := "sqlserver://" + user + ":" + pass + "@" + domain + ":" + port + "?database=" + database
// 	return dsn
// }

package config

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var KDB *gorm.DB

func init() {
	domain := "10.100.0.18"
	user := "ASISTENCIA"
	pass := "*T3cn0l0g14-*"
	port := "1433"
	database := "KACTUS"

	dsn := "sqlserver://" + user + ":" + pass + "@" + domain + ":" + port + "?database=" + database
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	KDB = db
	log.Println("Conexión a Kactus con éxito")
}
