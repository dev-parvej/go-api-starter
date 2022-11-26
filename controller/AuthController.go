package controller

import (
	"net/http"
	"time"

	"github.com/dev-parvej/go-api-starter-sql/config"
	"github.com/dev-parvej/go-api-starter-sql/db"
	"github.com/dev-parvej/go-api-starter-sql/dto"
	"github.com/dev-parvej/go-api-starter-sql/models"
	"github.com/dev-parvej/go-api-starter-sql/util"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var loginDto dto.LoginDto
	util.JsonDecoder(r, &loginDto)

	error := util.ValidateStruct(loginDto)

	if error != nil {
		util.Res.Status422().Writer(w).Data(error.Error())
		return
	}
	var user models.User

	db.Query().First(&user, "email=?", loginDto.Email)

	if user.IsEmpty() {
		util.Res.Status422().Writer(w).Data(map[string]string{"message": "userWithEmailNotFound"})
		return
	}

	if !util.ComparePassword(user.Password, loginDto.Password) {
		util.Res.Status403().Writer(w).Data(map[string]string{"message": "inCorrectEmailPassword"})
		return
	}

	var JWT_SECRET = []byte(config.Get("JWT_SECRET"))

	accessToken, accessTokenErr := generateAccessToken(user, JWT_SECRET)
	refreshToken, refreshTokenErr := generateRefreshToken(JWT_SECRET)

	if accessTokenErr != nil || refreshTokenErr != nil {
		util.Res.Writer(w).Status500().Data(map[string]any{"message": accessTokenErr.Error()})
		return
	}

	go storeRefreshToken(refreshToken, user.ID)

	util.Res.Writer(w).Status().Data(map[string]any{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"user":         user,
	})
}

func storeRefreshToken(token string, userId uint) {
	refreshToken := models.RefreshToken{
		RefreshToken: token,
		UserId:       userId,
	}

	db.Query().Create(&refreshToken)
}

func generateAccessToken(user models.User, JWT_SECRET []byte) (string, error) {
	claims := make(map[string]interface{})
	claims["exp"] = time.Now().Add(time.Duration(util.ParseInt(config.Get("ACCESS_TOKEN_EXPIRATION"))) * time.Minute)
	claims["user"] = user.ID

	return util.Token().CreateToken(claims)
}

func generateRefreshToken(JWT_SECRET []byte) (string, error) {
	claims := make(map[string]interface{})
	claims["exp"] = time.Now().Add(time.Duration(util.ParseInt(config.Get("REFRESH_TOKEN_EXPIRATION"))) * (time.Hour * 24))

	return util.Token().CreateToken(claims)
}
