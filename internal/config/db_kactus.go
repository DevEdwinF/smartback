package config

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var KDB *gorm.DB

func KactusDB() {
	domain := "DB-SEVEN-KACTUS"
	user := "ASISTENCIA"
	pass := "*T3cn0l0g14-+*"
	port := "9930"
	database := "gorm"

	dsn := "sqlserver://" + user + ":" + pass + "@" + domain + ":" + port + "?database=" + database
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	KDB = db
	log.Println("Conexión a Kactus con éxito")
}
