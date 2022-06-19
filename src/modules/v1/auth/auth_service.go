package auth

import (
	"fmt"

	"github.com/fajarihsan21/go-backend/src/database/gorm/models"
	help "github.com/fajarihsan21/go-backend/src/helpers"
	"github.com/fajarihsan21/go-backend/src/interfaces"
)

type token_res struct {
	Tokens string `json:"token"`
}

type auth_service struct {
	rep interfaces.UserRepo
}

func NewService(rep interfaces.UserRepo) *auth_service {
	return &auth_service{rep}
}

func (r *auth_service) Login(body *models.User) (*help.Response, error) {
	fmt.Println(body)
	user, err := r.rep.FindEmail(body.Email)
	if err != nil {
		return nil, err
	}

	if !help.CheckPassword(user.Password, body.Password) {
		return help.ResJSON(401, "wrong password"), nil
	}

	tkn := help.NewToken(body.Email)
	token, err := tkn.CreateToken()
	if err != nil {
		return nil, err
	}
	res := help.ResJSON(200, token_res{Tokens: token})
	return res, nil
}
