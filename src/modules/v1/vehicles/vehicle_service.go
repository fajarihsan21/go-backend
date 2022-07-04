package vehicles

import (
	"github.com/fajarihsan21/go-backend/src/database/gorm/models"
	help "github.com/fajarihsan21/go-backend/src/helpers"
	"github.com/fajarihsan21/go-backend/src/interfaces"
)

type vehicle_service struct {
	rep interfaces.VhcRepo
}

func NewService(rep interfaces.VhcRepo) *vehicle_service {
	return &vehicle_service{rep}
}

func (re *vehicle_service) FindVh() (*help.Response, error) {
	data, err := re.rep.FindVh()
	if err != nil {
		return nil, err
	}
	res := help.ResJSON(200, data)
	return res, nil
}

func (re *vehicle_service) FindIdVhc(id uint64) (*help.Response, error) {
	data, err := re.rep.FindIdVhc(id)
	if err != nil {
		return nil, err
	}

	res := help.ResJSON(200, data)
	return res, nil
}

func (re *vehicle_service) FindName(name string) (*help.Response, error) {
	data, err := re.rep.FindName(name)
	if err != nil {
		return nil, err
	}

	res := help.ResJSON(200, data)
	return res, nil
}

func (re *vehicle_service) Add(data *models.Vehicle) (*help.Response, error) {
	data, err := re.rep.Add(data)
	if err != nil {
		return nil, err
	}

	res := help.ResJSON(200, data)
	return res, nil
}

func (re *vehicle_service) Delete(id uint64) (*help.Response, error) {
	data, err := re.rep.FindIdVhc(id)
	if err != nil {
		return nil, err
	}

	err = re.rep.Delete(data.Id_vehicle)
	if err != nil {
		return nil, err
	}
	res := help.ResJSON(200, data)
	return res, nil
}

func (re *vehicle_service) Update(id uint64, data *models.Vehicle) (*help.Response, error) {
	data, err := re.rep.Update(id, data)
	if err != nil {
		return nil, err
	}

	res := help.ResJSON(200, data)
	return res, nil
}
