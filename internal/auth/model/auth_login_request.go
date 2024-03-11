package model

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
