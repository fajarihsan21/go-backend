package auth

import (
	"github.com/fajarihsan21/go-backend/src/modules/v1/users"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// localhost:8080/auth/ {post}
func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/auth/").Subrouter()

	repo := users.NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.SignIn).Methods("POST")
}
