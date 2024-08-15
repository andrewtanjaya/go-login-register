package routes

import (
	"go-login-register/controllers"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()

	router.HandleFunc("/register", controllers.SignUp).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
}
