package middleware

import (
	"net/http"
	"strings"

	help "github.com/fajarihsan21/go-backend/src/helpers"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			http.Error(w, "Invalid Header Type", http.StatusBadRequest)
			// help.ResJSON(401, "Invalid Header Type").Send(w)
			return
		}

		token := strings.Replace(headerToken, "Bearer ", "", -1)
		cToken, err := help.CheckToken(token)
		if err != nil {
			help.ResJSON(401, err.Error()).Send(w)
			return
		}

		if !cToken {
			help.ResJSON(401, "Please Relogin").Send(w)
			return
		}
		next.ServeHTTP(w, r)
	}
}
