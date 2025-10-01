package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) CreateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	s, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return s, nil
}
