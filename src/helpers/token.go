package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var mySecretKeys = []byte(os.Getenv("JWT_KEYS"))

type claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewToken(email string) *claims {
	return &claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
}

func (c *claims) CreateToken() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tokens.SignedString(mySecretKeys)
}

func CheckToken(tkn string) (bool, error) {
	tokens, err := jwt.ParseWithClaims(tkn, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecretKeys), nil
	})
	if err != nil {
		return false, err
	}
	return tokens.Valid, nil
}
