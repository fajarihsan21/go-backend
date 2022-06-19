package vehicles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fajarihsan21/go-backend/src/database/gorm/models"
	"github.com/fajarihsan21/go-backend/src/helpers"

	"github.com/gorilla/mux"
)

type vehicle_ctrl struct {
	repo *vehicle_repo
}

func NewCtrl(rep *vehicle_repo) *vehicle_ctrl {
	return &vehicle_ctrl{rep}
}

// GET ALL DATA
func (rep *vehicle_ctrl) GetAllVhcl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := rep.repo.FindVh()

	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	helpers.JSON(w, http.StatusOK, data)
}

// CREATE DATA
func (rep *vehicle_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.Vehicles
	json.NewDecoder(r.Body).Decode(&data)

	err := helpers.Validate(data)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	result, err := rep.repo.Add(&data)
	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
	}

	helpers.JSON(w, http.StatusCreated, &result)
}

// DELETE DATA
func (rep *vehicle_ctrl) DeleteData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data = mux.Vars(r)
	id, _ := strconv.Atoi(data["id_vehicle"])

	result, err := rep.repo.Delete(&id)

	if err != nil {
		helpers.ERROR(w, http.StatusInternalServerError, err)
	}

	helpers.JSON(w, http.StatusOK, &result)
}

// UPDATE DATA
func (rep *vehicle_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vhl *models.Vehicles
	var data = mux.Vars(r)["id_vehicle"]

	json.NewDecoder(r.Body).Decode(&vhl)

	err := helpers.Validate(vhl)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	id, _ := strconv.Atoi(data)
	result, err := rep.repo.Update(&id, vhl)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}
