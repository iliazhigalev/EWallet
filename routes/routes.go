package routes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// InitRoutes initializes the routes and injects the database connection.
func InitRoutes(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	return router
}
