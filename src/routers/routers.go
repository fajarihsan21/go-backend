package routers

import (
	"net/http"

	database "github.com/fajarihsan21/go-backend/src/database/gorm"
	"github.com/fajarihsan21/go-backend/src/modules/v1/auth"
	"github.com/fajarihsan21/go-backend/src/modules/v1/users"
	"github.com/fajarihsan21/go-backend/src/modules/v1/vehicles"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()
	nRelic, err := newrelic.NewApplication(
		newrelic.ConfigAppName("go-backend"),
		newrelic.ConfigLicense("4c1ab9865ed2ea01e08b3127a944df4ed5f4NRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		return nil, err
	}
	mainRoute.Use(nrgorilla.Middleware(nRelic))

	db, err := database.New()
	if err != nil {
		return nil, err
	}

	mainRoute.HandleFunc(newrelic.WrapHandleFunc(nRelic, "/relic", relicHandler)).Methods("GET")
	mainRoute.HandleFunc("/", rootHandler).Methods("GET")
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
func relicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello newrelic"))
}
