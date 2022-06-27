package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fajarihsan21/go-backend/src/database/gorm/models"
	"github.com/fajarihsan21/go-backend/src/helpers"
	"github.com/fajarihsan21/go-backend/src/interfaces"

	"github.com/gorilla/mux"
)

type user_ctrl struct {
	repo interfaces.UserService
}

func NewCtrl(rep interfaces.UserService) *user_ctrl {
	return &user_ctrl{rep}
}

// GET ALL DATA
func (rep *user_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := rep.repo.FindAll()

	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	data.Send(w)
}

// GET ID DATA
func (rep *user_ctrl) GetId(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)["id_user"]
	id, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	result, err := rep.repo.FindId(id)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
	}
	result.Send(w)
}

// GET EMAIL
func (rep *user_ctrl) FindEmail(w http.ResponseWriter, r *http.Request) {
	var data models.User
	json.NewDecoder(r.Body).Decode(&data)

	result, err := rep.repo.FindEmail(data.Email)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
	}
	result.Send(w)
}

// CREATE DATA or Register
func (rep *user_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = helpers.Validate(data)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	result, err := rep.repo.Add(&data)
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	result.Send(w)
}

// DELETE DATA
func (rep *user_ctrl) DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data = mux.Vars(r)["id_user"]
	id, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	result, err := rep.repo.Delete(id)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	result.Send(w)
}

// UPDATE DATA
func (rep *user_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data = r.URL.Query()
	var users models.User

	json.NewDecoder(r.Body).Decode(&users)

	err := helpers.Validate(users)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	id, err := strconv.ParseUint(data["id_user"][0], 10, 64)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
	}
	res, err := rep.repo.Update(id, &users)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
	}
	res.Send(w)
}
