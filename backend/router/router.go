package router

import (
	"../controllers"
	"github.com/gorilla/mux"
)

// Router is the main switch for all routes used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/task", controllers.GetAllTask).Methods("GET", "OPTIONS")

	return router
}
