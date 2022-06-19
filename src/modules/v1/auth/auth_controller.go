package auth

import (
	"encoding/json"
	"net/http"

	"github.com/fajarihsan21/go-backend/src/database/gorm/models"
	"github.com/fajarihsan21/go-backend/src/helpers"
	"github.com/fajarihsan21/go-backend/src/interfaces"
)

type auth_ctrl struct {
	rep interfaces.AuthService
}

func NewCtrl(rep interfaces.AuthService) *auth_ctrl {
	return &auth_ctrl{rep}
}

func (re *auth_ctrl) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	result, err := re.rep.Login(&data)

	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	result.Send(w)
}
