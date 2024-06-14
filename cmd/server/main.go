// главный пакет, с него начинается вход в программу
package main

import (
	"ewallet/config"
	"ewallet/internal/database"
	"ewallet/routes"
	"log"
	"net/http"
)

func main() {

	config.InitConfig()

	_, err := database.ConnectDb()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	router := routes.InitRoutes()

	log.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
