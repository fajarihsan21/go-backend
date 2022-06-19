package interfaces

import (
	"github.com/fajarihsan21/go-backend/src/database/gorm/models"
	help "github.com/fajarihsan21/go-backend/src/helpers"
)

type AuthService interface {
	Login(body *models.User) (*help.Response, error)
}
