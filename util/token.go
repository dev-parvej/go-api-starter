package util

import (
	"github.com/dev-parvej/go-api-starter-sql/config"
	"github.com/golang-jwt/jwt"
)

type JWTToken struct {
	Secret []byte
}

func Token() *JWTToken {
	return &JWTToken{
		Secret: []byte(config.Get("JWT_SECRET")),
	}
}

func (jwtToken *JWTToken) CreateToken(payload map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims = CopyMap(payload, claims)

	tokenString, err := token.SignedString(jwtToken.Secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
