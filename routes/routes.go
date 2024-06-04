package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// InitRoutes initializes the routes and injects the database connection.
func InitRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	// Регистрация обработчика для маршрута "/hello"
	router.HandleFunc("/hello", helloHandler)

	return router
}

// Обработчик для маршрута "/hello"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Отправка приветственного сообщения
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Привет, мир"))
}
