package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID        uuid.UUID `gorm:"column:id;type:varchar(36);primary_key" json:"id"`
	UserID    string    `gorm:"column:user_id;type:varchar(36);index;uniqueIndex:idx_accounts_balance_currency" json:"user_id"`
	Balance   int64     `gorm:"column:balance" json:"balance"`
	Currency  string    `gorm:"column:currency;type:varchar(50);uniqueIndex:idx_accounts_balance_currency" json:"currency"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (a *Account) TableName() string {
	return "accounts"
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	tx.Statement.SetColumn("id", uuid)
	return nil
}
