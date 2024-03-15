package entity

import (
	"time"
)

type User struct {
	ID        string    `gorm:"column:id;primaryKey;size:36"`
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
