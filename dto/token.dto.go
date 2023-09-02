package dto

type TokenDto struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
