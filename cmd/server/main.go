package main

import (
	"log"
	"net/http"

	"ewallet/internal/config"
	"ewallet/internal/database"
	"ewallet/internal/handler"
	"ewallet/internal/repository/transaction"
	"ewallet/internal/repository/wallet"
	"ewallet/internal/routes"
	"ewallet/internal/service"
)

func main() {

	config.InitConfig()

	walletRepo := wallet.New()
	transactionRepo := transaction.New()
	db, err := database.ConnectDb()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	eWalletService := service.NewEWallet(db, walletRepo, transactionRepo)
	handlerItem := handler.New(eWalletService)

	router := routes.InitRoutes(handlerItem)

	log.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
