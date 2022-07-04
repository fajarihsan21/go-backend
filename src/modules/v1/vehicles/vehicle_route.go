package vehicles

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// localhost:8080/vehicles/ {get, post}
func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicles/").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	// route.HandleFunc("/", middleware.Do(ctrl.GetAllVhcl, middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/", ctrl.GetAllVhcl).Methods("GET")
	route.HandleFunc("/", ctrl.AddData).Methods("POST")
	route.HandleFunc("/{id_vehicle}", ctrl.DeleteData).Methods("DELETE")
	route.HandleFunc("/update", ctrl.Update).Methods("PUT")

}
