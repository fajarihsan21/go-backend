package vehicles

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fajarihsan21/go-backend/src/database/gorm/models"
	"github.com/fajarihsan21/go-backend/src/helpers"
	"github.com/fajarihsan21/go-backend/src/interfaces"

	"github.com/gorilla/mux"
)

type vehicle_ctrl struct {
	repo interfaces.VhcService
}

func NewCtrl(rep interfaces.VhcService) *vehicle_ctrl {
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

// GET BY ID
func (rep *vehicle_ctrl) GetId(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)["id_vehicle"]
	id, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	result, err := rep.repo.FindIdVhc(id)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
	}
	result.Send(w)
}

// GET Name
func (rep *vehicle_ctrl) FindName(w http.ResponseWriter, r *http.Request) {
	var data models.Vehicle
	json.NewDecoder(r.Body).Decode(&data)

	result, err := rep.repo.FindName(data.VehicleName)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
	}
	result.Send(w)

}

// GET Category
func (rep *vehicle_ctrl) FindCategory(w http.ResponseWriter, r *http.Request) {
	var data models.Vehicle
	json.NewDecoder(r.Body).Decode(&data)

	result, err := rep.repo.FindCategory(data.Category)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
	}
	result.Send(w)
}

// CREATE DATA
func (rep *vehicle_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.Vehicle
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

	var data = mux.Vars(r)["id_vehicle"]
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
func (rep *vehicle_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data = r.URL.Query()
	var vhc models.Vehicle

	json.NewDecoder(r.Body).Decode(&vhc)

	err := helpers.Validate(vhc)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
		return
	}

	id, err := strconv.ParseUint(data["id_vehicle"][0], 10, 64)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
	}
	res, err := rep.repo.Update(id, &vhc)
	if err != nil {
		helpers.ERROR(w, http.StatusBadRequest, err)
	}
	res.Send(w)
}
