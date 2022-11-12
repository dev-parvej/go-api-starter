package controller

import (
	"log"
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
		log.Default().Println(errors.Error())
		util.JsonEncoder(w, errors.Error())
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

	util.JsonEncoder(w, user)
}
