package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Product struct {
	ID        uuid.UUID      `gorm:"column:id;primaryKey;size:36"`
	Name      string         `gorm:"column:name;size:191;not null"`
	Stock     int            `gorm:"column:stock;default:0"`
	Images    datatypes.JSON `gorm:"type:text[];not null"`
	UserID    uuid.UUID      `gorm:"size:36"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
}

// TableName overrides the table name used
func (u *Product) TableName() string {
	return "products"
}

func (u *Product) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	tx.Statement.SetColumn("id", uuid)
	return nil
}
