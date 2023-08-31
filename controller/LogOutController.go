package controller

import (
	"net/http"

	"github.com/dev-parvej/go-api-starter-sql/db"
	"github.com/dev-parvej/go-api-starter-sql/models"
	"github.com/dev-parvej/go-api-starter-sql/util"
)

type LogOutDto struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	logOutDto, error := util.ValidateRequest(r, LogOutDto{})

	if error != nil {
		util.Res.Status403().Writer(w).Data(error.Error())
		return
	}

	db.Query().Where("refresh_token = ?", logOutDto.RefreshToken).Delete(models.RefreshToken{})

	util.Res.Writer(w).Status().Data(map[string]string{"message": "LogOutSuccessful"})
}
