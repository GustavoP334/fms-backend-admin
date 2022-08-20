package main

import (
	DB "application-web/DB"
	"application-web/models"
	"application-web/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DB.DbConnection()
	DB.Conn.AutoMigrate(&models.Driver{}, &models.Address{}, &models.FieldsDocuments{})
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
