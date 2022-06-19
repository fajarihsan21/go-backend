package vehicles

import (
	"github.com/fajarihsan21/go-backend/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// localhost:8080/vehicles/ {get, post}
func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicles/").Subrouter()

	repo := NewRepo(db)
	ctrl := NewCtrl(repo)

	route.HandleFunc("/", middleware.Do(ctrl.GetAllVhcl, middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/", ctrl.AddData).Methods("POST")
	route.HandleFunc("/{id_user}", ctrl.DeleteData).Methods("DELETE")
	route.HandleFunc("/{id_user}", ctrl.Update).Methods("PUT")

}
