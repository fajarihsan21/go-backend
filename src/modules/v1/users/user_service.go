package users

import (
	"github.com/fajarihsan21/go-backend/src/database/gorm/models"
	help "github.com/fajarihsan21/go-backend/src/helpers"
	"github.com/fajarihsan21/go-backend/src/interfaces"
)

type user_service struct {
	rep interfaces.UserRepo
}

func NewService(rep interfaces.UserRepo) *user_service {
	return &user_service{rep}
}

func (re *user_service) FindAll() (*help.Response, error) {
	data, err := re.rep.FindAll()
	if err != nil {
		return nil, err
	}
	res := help.ResJSON(200, data)
	return res, nil
}

func (re *user_service) FindId(id uint64) (*help.Response, error) {
	data, err := re.rep.FindId(id)
	if err != nil {
		return nil, err
	}

	res := help.ResJSON(200, data)
	return res, nil
}

func (re *user_service) FindEmail(email string) (*help.Response, error) {
	data, err := re.rep.FindEmail(email)
	if err != nil {
		return nil, err
	}

	res := help.ResJSON(200, data)
	return res, nil
}

func (re *user_service) Add(usr *models.User) (*help.Response, error) {
	hashPwd, err := help.HashPassword(usr.Password)
	if err != nil {
		return nil, err
	}

	usr.Password = hashPwd
	data, err := re.rep.Add(usr)
	if err != nil {
		return nil, err
	}

	res := help.ResJSON(200, data)
	return res, nil
}

func (re *user_service) Delete(id uint64) (*help.Response, error) {
	data, err := re.rep.FindId(id)
	if err != nil {
		return nil, err
	}

	err = re.rep.Delete(data.IdUser)
	if err != nil {
		return nil, err
	}
	res := help.ResJSON(200, data)
	return res, nil
}

func (re *user_service) Update(id uint64, data *models.User) (*help.Response, error) {
	data, err := re.rep.Update(id, data)
	if err != nil {
		return nil, err
	}

	res := help.ResJSON(200, data)
	return res, nil
}
