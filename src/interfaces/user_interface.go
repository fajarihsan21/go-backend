package interfaces

import (
	"github.com/fajarihsan21/go-backend/src/database/gorm/models"
	help "github.com/fajarihsan21/go-backend/src/helpers"
)

type UserRepo interface {
	FindAll() (*models.Users, error)
	FindId(id uint64) (*models.User, error)
	FindEmail(email string) (*models.User, error)
	Add(*models.User) (*models.User, error)
	Delete(id uint) error
	Update(id uint64, data *models.User) (*models.User, error)
}
type UserService interface {
	FindAll() (*help.Response, error)
	FindId(id uint64) (*help.Response, error)
	FindEmail(email string) (*help.Response, error)
	Add(*models.User) (*help.Response, error)
	Delete(id uint64) (*help.Response, error)
	Update(id uint64, data *models.User) (*help.Response, error)
}
