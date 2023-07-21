package config

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var KDB *gorm.DB

func init() {
	domain := "DB-SEVEN-KACTUS"
	user := "CAPACITA"
	pass := "12345678"
	port := "1433"
	database := "gorm"

	dsn := "sqlserver://" + user + ":" + pass + "@" + domain + ":" + port + "?database=" + database
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	KDB = db
	log.Println("Conexión a Kactus con éxito")
}
