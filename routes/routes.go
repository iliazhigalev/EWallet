package routes

import (
	"ewallet/pkg/handler"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// InitRoutes initializes the routes and injects the database connection.
func InitRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	// Регистрация обработчика для маршрута "/hello"
	router.HandleFunc("/hello", handler.HelloHandler).Methods("GET")

	return router
}
