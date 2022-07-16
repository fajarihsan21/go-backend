package users

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// localhost:8080/users/ {get, post}
func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/users/").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/{id_user}", ctrl.GetId).Methods("GET")
	route.HandleFunc("/", ctrl.AddData).Methods("POST")
	route.HandleFunc("/{id_user}", ctrl.DeleteData).Methods("DELETE")
	route.HandleFunc("/update", ctrl.Update).Methods("PUT")
}
