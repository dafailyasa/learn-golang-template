package model

import (
	"time"

	"github.com/google/uuid"
)

type CreateAccountResponse struct {
	ID        uuid.UUID `json:"id"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
}
