package model

type AuthRegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=5"`
	Name     string `json:"name" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
}
