package models

type RefreshToken struct {
	ID           uint   `json:"id" gorm:"column:id;primarykey"`
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token;type:varchar(256)"`
	UserId       uint   `json:"user_id" gorm:"column:user_id"`
	User         User
}
