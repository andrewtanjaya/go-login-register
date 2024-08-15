package main

import (
	"go-login-register/configs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDatabase()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
