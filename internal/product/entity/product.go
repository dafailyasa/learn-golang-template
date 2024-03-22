package entity

import (
	"time"

	"github.com/dafailyasa/learn-golang-template/internal/auth/entity"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID      `json:"id" gorm:"column:id;primaryKey;size:36"`
	Name        string         `json:"name" gorm:"column:name;size:191;notnull"`
	Stock       uint32         `json:"stock" gorm:"column:stock;notnull;default:0"`
	Images      datatypes.JSON `json:"images" gorm:"type:text[];notnull"`
	Price       float64        `json:"price" gorm:"column:price;type:bigint; notnull"`
	IsPublished bool           `json:"isPublished" gorm:"column:isPublished;default:false;type:bool;comment:0 = false 1 = true"`
	UserID      uuid.UUID      `json:"userId" gorm:"size:36"`
	User        entity.User    `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time      `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   time.Time      `json:"updatedAt" gorm:"column:updated_at"`
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
