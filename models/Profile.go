package models

type Profile struct {
	ID     uint `json:"id" gorm:"column=id,primaryKey,autoIncrement"`
	UserId uint `json:"user_id" gorm:"column=user_id"`
	User   User
}
