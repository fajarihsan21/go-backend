package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Vehicles struct {
	Name     string
	Category string
	Price    int
}

func handlerReqs() {
	mainRoute := mux.NewRouter()
	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")
	mainRoute.HandleFunc("/users/{id}", paramsHandler).Methods("GET")
	mainRoute.HandleFunc("/vehicles", queryHandler).Methods("GET")
	mainRoute.HandleFunc("/vehicles", bodyHandler).Methods("POST")

	if err := http.ListenAndServe("localhost:8080", mainRoute); err != nil {
		log.Fatal("Application failed to run")
	}
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello backend"))
}

func paramsHandler(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	fmt.Fprintf(w, "Params: %v", vars["id"])
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	var vars = r.URL.Query()
	fmt.Fprintf(w, "Params: %v", vars["name"][0])
}

func bodyHandler(w http.ResponseWriter, r *http.Request) {
	var rentV Vehicles
	json.NewDecoder(r.Body).Decode(&rentV)
	fmt.Println(rentV.Name)
	// fmt.Fprintf(w, "Params: %v", vars["name"][0])
}
