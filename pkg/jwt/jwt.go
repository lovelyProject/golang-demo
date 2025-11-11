package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	Email string `json:"email"`
	ID    uint   `json:"id"`
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) CreateToken(data JWTData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
		"id":    data.ID,
	})

	s, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return s, nil
}

func (j *JWT) Parse(token string) (bool, *JWTData) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}

	fmt.Println(t)
	email := t.Claims.(jwt.MapClaims)["email"]
	idRaw := t.Claims.(jwt.MapClaims)["id"]
	if idRaw == nil {
		{
			return false, nil
		}
	}
	var idUint uint
	if f, ok := idRaw.(float64); ok {
		idUint = uint(f)
	}
	return t.Valid, &JWTData{Email: email.(string), ID: idUint}
}
