package DB

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func DbConnection() *gorm.DB {
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
	dbname := os.Getenv("DBNAME")
	sslmode := os.Getenv("DBSSLMODE")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Conn = db

	return db
}
