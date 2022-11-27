package dto

type RefreshDto struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
