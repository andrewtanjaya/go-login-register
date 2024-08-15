package routes

import (
	"go-login-register/controllers"
	"go-login-register/middlewares"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()

	router.Use(middlewares.Auth)
	router.HandleFunc("/me", controllers.Me).Methods("GET")
}
