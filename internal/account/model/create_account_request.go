package model

type CreateAccountRequest struct {
	Currency string `json:"currency" validate:"required"`
}
