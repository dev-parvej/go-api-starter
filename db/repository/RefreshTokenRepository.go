package repository

import (
	"time"

	"github.com/dev-parvej/go-api-starter-sql/config"
	"github.com/dev-parvej/go-api-starter-sql/db"
	"github.com/dev-parvej/go-api-starter-sql/models"
	"github.com/dev-parvej/go-api-starter-sql/util"
)

type rTRepository struct{}

func (r *rTRepository) Insert(token string, userId int) {
	refreshToken := models.RefreshToken{
		RefreshToken: token,
		UserId:       userId,
		ValidUntil:   time.Now().Add(time.Duration(util.ParseInt(config.Get("REFRESH_TOKEN_EXPIRATION"))) * (time.Hour * 24)),
	}

	db.Query().Create(&refreshToken)
}

func (r *rTRepository) Delete(refreshToken string) {
	db.Query().Where("refresh_token = ?", refreshToken).Delete(&models.RefreshToken{})
}

var RefreshTokenRepository = &rTRepository{}
