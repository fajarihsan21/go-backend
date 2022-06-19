package routers

import (
	"net/http"

	database "github.com/fajarihsan21/go-backend/src/database/gorm"
	"github.com/fajarihsan21/go-backend/src/modules/v1/auth"
	"github.com/fajarihsan21/go-backend/src/modules/v1/users"
	"github.com/fajarihsan21/go-backend/src/modules/v1/vehicles"
	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()
	mainRoute.HandleFunc("/", rootHandler).Methods("GET")

	db, err := database.New()
	if err != nil {
		return nil, err
	}

	// localhost:8080/users/
	users.New(mainRoute, db)
	// localhost:8080/vehicles/
	vehicles.New(mainRoute, db)
	// localhost:8080/auth/
	auth.New(mainRoute, db)

	return mainRoute, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello backend"))
}
