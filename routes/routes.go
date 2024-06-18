package routes

import (
	"ewallet/internal/handler"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/wallet", handler.CreateWallet).Methods("POST")
	router.HandleFunc("/api/v1/wallet/{walletId}/send", handler.SendMoney).Methods("POST")
	router.HandleFunc("/api/v1/wallet/{walletId}/history", handler.GetTransactionHistory).Methods("GET")
	router.HandleFunc("/api/v1/wallet/{walletId}", handler.GetWallet).Methods("GET")

	// Маршрут для отображения Swagger-документации
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/openapi.yaml"), // Ссылка на OpenAPI файл
	))

	// Маршрут для отдачи OpenAPI yaml файла
	router.HandleFunc("/swagger/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("api", "openapi.yaml"))
	})

	return router
}
