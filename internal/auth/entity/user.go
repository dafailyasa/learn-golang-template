package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey;size:36"`
	Password  string    `gorm:"column:password"`
	Email     string    `gorm:"column:email;uniqueIndex;size:191"`
	Name      string    `gorm:"column:name;size:191"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName overrides the table name used
func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	tx.Statement.SetColumn("id", uuid)
	return nil
}
