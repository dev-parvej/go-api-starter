package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/dev-parvej/go-api-starter-sql/db"
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

	_, accessTokenErr := util.Token().VerifyToken(tokenDto.AccessToken)

	if errors.Is(accessTokenErr, util.ErrInvalidToken) {
		util.Res.Writer(w).Status403().Data(map[string]string{"message": "invalidToken"})
		return
	}

	refreshToken := models.RefreshToken{}

	db.Query().First(&refreshToken, "refresh_token=? AND DATE(valid_until) >= ?", tokenDto.RefreshToken, time.Now())

	if refreshToken.ID == 0 {
		util.Res.Writer(w).Status403().Data(map[string]string{"message": "tokenHasBeenRevoked"})
		return
	}

	util.Res.Writer(w).Status().Data(refreshToken)

}
