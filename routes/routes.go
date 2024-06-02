// routes/routes.go
package routes

import (
	"github.com/SamiranDas2004/go-auth/controller"
	"github.com/gorilla/mux"
	// "net/http"
)

// Setup returns the router with defined routes
func Setup() *mux.Router {
	// Define your routes here
	router := mux.NewRouter()
	//router.HandleFunc("/", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/registers", controller.Register).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	// router.HandleFunc("/", controller.Hello).Methods("GET")
	return router
}
