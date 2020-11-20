package router

import (
	"../controllers"

	"github.com/gorilla/mux"
)

// Router is the main switch for all routes used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/task", controllers.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", controllers.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", controllers.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTasks", controllers.DeleteAllTask).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/signin", controllers.SignIn).Methods("POST", "OPTIONS")
	router.HandleFunc("/register", controllers.CreateUser).Methods("POST", "OPTIONS")

	return router
}
