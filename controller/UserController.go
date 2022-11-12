package controller

import (
	"net/http"

	"github.com/dev-parvej/go-api-starter-sql/db"
	"github.com/dev-parvej/go-api-starter-sql/dto"
	"github.com/dev-parvej/go-api-starter-sql/models"
	"github.com/dev-parvej/go-api-starter-sql/util"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserDto dto.CreateUserDto

	util.JsonDecoder(r, &createUserDto)

	errors := util.ValidateStruct(createUserDto)

	if errors != nil {
		util.Res.Writer(w).Status422().Data(errors.Error())
		return
	}

	var existingUser models.User

	db.DB().First(&existingUser, "email=?", createUserDto.Email)

	if !existingUser.IsEmpty() {
		util.Res.Writer(w).Status422().Data(map[string]any{"message": "UniqueEmail"})
		return
	}

	password, _ := util.HashPassword(createUserDto.Password)

	user := models.User{
		FirstName: createUserDto.FirstName,
		LastName:  createUserDto.LastName,
		Email:     createUserDto.Email,
		Password:  password,
	}

	db.Model(models.User{}).Create(&user)

	util.Res.Writer(w).Status().Data(user)
}
