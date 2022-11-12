package models

type User struct {
	BaseModel
	FirstName string `json:"first_name" gorm:"column:first_name;type:varchar(255)"`
	LastName  string `json:"last_name" gorm:"column:last_name;type:varchar(255)"`
	Email     string `json:"email" gorm:"uniqueIndex;column:email;type:varchar(255)"`
	Password  string `json:"-" gorm:"column:password;type:varchar(255)"`
}
