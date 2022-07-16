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

func (r *vehicle_repo) FindVh() (*models.Vehicles, error) {
	var vehicles models.Vehicles
	result := r.db.Find(&vehicles)

	if result.Error != nil {
		return nil, errors.New("error getting all data")
	}
	return &vehicles, nil
}

func (r *vehicle_repo) FindIdVhc(id uint64) (*models.Vehicle, error) {
	var data models.Vehicle
	result := r.db.First(&data, "id_vehicle = ?", id)

	if result.Error != nil {
		return nil, errors.New("failed getting id")
	}
	return &data, nil
}

func (r *vehicle_repo) FindName(name string) (*models.Vehicle, error) {
	var data models.Vehicle
	result := r.db.First(&data, "vehicle_name = ?", name)

	if result.Error != nil {
		return nil, errors.New("failed getting Vehicle Name")
	}
	return &data, nil

}

func (r *vehicle_repo) FindCategory(category string) (*models.Vehicle, error) {
	var data models.Vehicle
	result := r.db.First(&data, "category = ?", category)

	if result.Error != nil {
		return nil, errors.New("failed getting Vehicle Name")
	}
	return &data, nil
}

func (r *vehicle_repo) Add(data *models.Vehicle) (*models.Vehicle, error) {
	result := r.db.Create(data)

	if result.Error != nil {
		return nil, errors.New("error saving the data")
	}
	return data, nil
}

// func (r *vehicle_repo) Delete(data *int) (*models.Vehicle, error) {
// 	var vehicles models.Vehicles
// 	r.db.First(&vehicles, &data)
// 	result := r.db.Delete(&models.Vehicle{}, &data)

// 	if result.Error != nil {
// 		return nil, errors.New("error deleting the data")
// 	}
// 	return &models.Vehicle{}, nil
// }

func (r *vehicle_repo) Delete(id uint) error {
	var data models.Vehicle
	result := r.db.Where("id_vehicle = ?", id).Delete(&data)
	if result.Error != nil {
		return errors.New("error deleting the data")
	}

	return nil
}

// func (r *vehicle_repo) Update(Id *int, data *models.Vehicle) (*models.Vehicle, error) {
// 	var vhcl models.Vehicle
// 	res := r.db.First(&vhcl, "id_vehicle = ?", Id)
// 	if res.Error != nil {
// 		return nil, errors.New("error getting ID")
// 	}

// 	result := r.db.Save(&data)
// 	if result.Error != nil {
// 		return nil, errors.New("error updating data")
// 	}
// 	return data, nil
// }

func (r *vehicle_repo) Update(id uint64, data *models.Vehicle) (*models.Vehicle, error) {
	var vhc models.Vehicle
	result := r.db.Model(&vhc).Where("id_vehicle = ?", id).Updates(data)
	if result.Error != nil {
		return nil, errors.New("error deleting data")
	}
	return data, nil
}
