package models

import "time"

type RefreshToken struct {
	ID           int       `json:"id" gorm:"column:id;primarykey"`
	RefreshToken string    `json:"refresh_token" gorm:"column:refresh_token;type:varchar(256)"`
	ValidUntil   time.Time `json:"valid_until" gorm:"column:valid_until;type:date"`
	UserId       int       `json:"user_id" gorm:"column:user_id"`
	User         User
}
