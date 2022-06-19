package vehicles

import (
	"errors"

	"github.com/fajarihsan21/go-backend/src/database/gorm/models"

	"gorm.io/gorm"
)

type vehicle_repo struct {
	db *gorm.DB
}

func NewRepo(grm *gorm.DB) *vehicle_repo {
	return &vehicle_repo{grm}
}

func (r *vehicle_repo) FindVh() (*models.Vehicle, error) {
	var vehicles models.Vehicle
	result := r.db.Find(&vehicles)

	if result.Error != nil {
		return nil, errors.New("error getting all data")
	}
	return &vehicles, nil
}

func (r *vehicle_repo) Add(data *models.Vehicles) (*models.Vehicles, error) {
	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("error saving the data")
	}
	return data, nil
}

func (r *vehicle_repo) Delete(data *int) (*models.Vehicles, error) {
	var vehicles models.Vehicle
	r.db.First(&vehicles, &data)
	result := r.db.Delete(&models.Vehicles{}, &data)

	if result.Error != nil {
		return nil, errors.New("error deleting the data")
	}
	return &models.Vehicles{}, nil
}

func (r *vehicle_repo) Update(Id *int, data *models.Vehicles) (*models.Vehicles, error) {
	var vhcl models.Vehicles
	res := r.db.First(&vhcl, "id_vehicle = ?", Id)
	if res.Error != nil {
		return nil, errors.New("error getting ID")
	}

	result := r.db.Save(&data)
	if result.Error != nil {
		return nil, errors.New("error updating data")
	}
	return data, nil
}
