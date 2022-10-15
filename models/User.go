package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"column=name,type=varchar,size=255"`
	Email    string `json:"email" gorm:"column=email,type=varchar,size=255,unique,not null"`
	Password string `json:"-" gorm:"column=password,type=varchar"`
}
