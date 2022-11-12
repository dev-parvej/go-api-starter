package dto

type CreateUserDto struct {
	FirstName string `validate:"required,lte=255" json:"first_name"`
	LastName  string `validate:"lte=255" json:"last_name"`
	Email     string `validate:"required,lte=255,email" json:"email"`
	Password  string `validate:"required,gte=6,lte=20" json:"password"`
}
