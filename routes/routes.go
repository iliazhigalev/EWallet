package routes

import (
	"ewallet/internal/handler"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/wallet", handler.CreateWallet).Methods("POST")
	router.HandleFunc("/api/v1/wallet/{walletId}/send", handler.SendMoney).Methods("POST")
	router.HandleFunc("/api/v1/wallet/{walletId}/history", handler.GetTransactionHistory).Methods("GET")
	router.HandleFunc("/api/v1/wallet/{walletId}", handler.GetWallet).Methods("GET")

	return router
}
