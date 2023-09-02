package controller

import (
	"net/http"
	"time"

	"github.com/dev-parvej/go-api-starter-sql/db"
	"github.com/dev-parvej/go-api-starter-sql/db/repository"
	"github.com/dev-parvej/go-api-starter-sql/dto"
	"github.com/dev-parvej/go-api-starter-sql/models"
	"github.com/dev-parvej/go-api-starter-sql/util"
)

func GrantAccessToken(w http.ResponseWriter, r *http.Request) {
	tokenDto, err := util.ValidateRequest(r, dto.TokenDto{})

	if err != nil {
		util.Res.Writer(w).Status403().Data(err.Error())
		return
	}

	_, refreshTokenErr := util.Token().VerifyToken(tokenDto.RefreshToken)

	if refreshTokenErr != nil {
		util.Res.Writer(w).Status403().Data(map[string]string{"message": "invalidToken"})
		return
	}

	refreshTokenEntity := models.RefreshToken{}

	db.Query().First(&refreshTokenEntity, "refresh_token=? AND DATE(valid_until) >= ?", tokenDto.RefreshToken, time.Now())

	if refreshTokenEntity.ID == 0 {
		util.Res.Writer(w).Status403().Data(map[string]string{"message": "tokenHasBeenRevoked"})
		return
	}

	refreshToken, rTokenErr := util.Token().RefreshToken()
	accessToken, aTokenErr := util.Token().RefreshToken()

	if rTokenErr != nil || aTokenErr != nil {
		util.Res.Status403().Writer(w).Data(map[string]string{"message": "tokenCanNotBeGenerated"})
		return
	}

	go repository.RefreshTokenRepository.Insert(refreshToken, refreshTokenEntity.UserId)

	go repository.RefreshTokenRepository.Delete(tokenDto.RefreshToken)

	util.Res.Writer(w).Status().Data(map[string]string{"accessToken": accessToken, "refreshToken": refreshToken})
}
