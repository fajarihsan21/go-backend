package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func subRouter() {
	mainRoute := mux.NewRouter()
	mainRoute.HandleFunc("/", homeHandler).Methods("GET")

	// localhost:8080/users {get, post}
	users := mainRoute.PathPrefix("/users/").Subrouter()
	users.HandleFunc("/", usersHandler)

	// localhost:8080/vehicles {get, post}
	vehicles := mainRoute.PathPrefix("/vehicles/").Subrouter()
	vehicles.HandleFunc("/", vehiclesHandler)

	if err := http.ListenAndServe("127.0.0.1:8080", mainRoute); err != nil {
		log.Fatal("Application failed to run")
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello backend"))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from users"))
}

func vehiclesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from vehicles"))
}
