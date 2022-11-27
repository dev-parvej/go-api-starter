package controller

import (
	"net/http"
	"time"

	"github.com/dev-parvej/go-api-starter-sql/db"
	"github.com/dev-parvej/go-api-starter-sql/dto"
	"github.com/dev-parvej/go-api-starter-sql/models"
	"github.com/dev-parvej/go-api-starter-sql/util"
)

func CreateRefreshToken(w http.ResponseWriter, r *http.Request) {
	refreshTokenDto := dto.RefreshDto{}
	util.JsonDecoder(r, &refreshTokenDto)

	_, err := util.Token().VerifyToken(refreshTokenDto.RefreshToken)

	if err != nil {
		util.Res.Status403().Data(map[string]string{"message": "invalidToken"})
		return
	}

	refreshToken := models.RefreshToken{}

	db.Query().First(&refreshToken, "refresh_token=? AND DATE(valid_until) >= ? AND user_id=?", refreshTokenDto.RefreshToken, time.Now(), r.Header.Get("user_id"))

	if refreshToken.ID == 0 {
		util.Res.Status403().Data(map[string]string{"message": "tokenHasBeenRevoked"})
		return
	}

}
