package routes

import (
	"net/http"

	"ewallet/internal/handler"

	"github.com/gorilla/mux"
)

type handlerItem interface {
	CreateWallet(w http.ResponseWriter, r *http.Request)
	SendMoney(w http.ResponseWriter, r *http.Request)
	GetTransactionHistory(w http.ResponseWriter, r *http.Request)
	GetWallet(w http.ResponseWriter, r *http.Request)
}

func InitRoutes(handlerItem handlerItem) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/wallet", handler.CreateWallet).Methods("POST")
	router.HandleFunc("/api/v1/wallet/{walletId}/send", handler.SendMoney).Methods("POST")
	router.HandleFunc("/api/v1/wallet/{walletId}/history", handler.GetTransactionHistory).Methods("GET")
	router.HandleFunc("/api/v1/wallet/{walletId}", handler.GetWallet).Methods("GET")

	router.HandleFunc("/api/v2/wallet", handlerItem.CreateWallet).Methods("POST")
	router.HandleFunc("/api/v2/wallet/{walletId}/send", handlerItem.SendMoney).Methods("POST")
	router.HandleFunc("/api/v2/wallet/{walletId}/history", handlerItem.GetTransactionHistory).Methods("GET")
	router.HandleFunc("/api/v2/wallet/{walletId}", handlerItem.GetWallet).Methods("GET")

	return router
}
