package util

import (
	"errors"
	"time"

	"github.com/dev-parvej/go-api-starter-sql/config"
	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type JWTToken struct {
	secret []byte
}

func Token() *JWTToken {
	return &JWTToken{
		secret: []byte(config.Get("JWT_SECRET")),
	}
}

type Payload struct {
	UserId    int       `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func (jwtToken *JWTToken) AccessToken(userId int) (string, error) {
	claims := &Payload{
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Duration(ParseInt(config.Get("ACCESS_TOKEN_EXPIRATION"))) * time.Minute),
		UserId:    userId,
	}

	return jwtToken.createToken(claims)
}

func (jwtToken *JWTToken) RefreshToken() (string, error) {
	claims := &Payload{
		ExpiredAt: time.Now().Add(time.Duration(ParseInt(config.Get("REFRESH_TOKEN_EXPIRATION"))) * (time.Hour * 24)),
	}

	return jwtToken.createToken(claims)
}

func (jwtToken *JWTToken) createToken(payload *Payload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenString, err := token.SignedString(jwtToken.secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jwtToken *JWTToken) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(jwtToken.secret), nil
	}

	jwtClaims, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		err, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(err.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtClaims.Claims.(*Payload)

	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
