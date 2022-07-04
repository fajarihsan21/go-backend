package interfaces

import (
	"github.com/fajarihsan21/go-backend/src/database/gorm/models"
	help "github.com/fajarihsan21/go-backend/src/helpers"
)

type VhcRepo interface {
	FindVh() (*models.Vehicles, error)
	FindIdVhc(id uint64) (*models.Vehicle, error)
	FindName(name string) (*models.Vehicle, error)
	Add(data *models.Vehicle) (*models.Vehicle, error)
	Delete(id uint) error
	Update(id uint64, data *models.Vehicle) (*models.Vehicle, error)
}

type VhcService interface {
	FindVh() (*help.Response, error)
	FindIdVhc(id uint64) (*help.Response, error)
	FindName(name string) (*help.Response, error)
	Add(data *models.Vehicle) (*help.Response, error)
	Delete(id uint64) (*help.Response, error)
	Update(id uint64, data *models.Vehicle) (*help.Response, error)
}
