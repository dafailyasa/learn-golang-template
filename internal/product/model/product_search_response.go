package model

import (
	"time"

	"github.com/dafailyasa/learn-golang-template/internal/product/entity"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type userSearchProductResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type SearchProductResponse struct {
	ID        uuid.UUID                 `json:"id"`
	Name      string                    `json:"name"`
	Stock     uint32                    `json:"stock"`
	Images    datatypes.JSON            `json:"images"`
	Price     float64                   `json:"price"`
	User      userSearchProductResponse `json:"user"`
	CreatedAt time.Time                 `json:"createdAt"`
	UpdatedAt time.Time                 `json:"updatedAt"`
}

func ConvertProductsRes(products *[]entity.Product) []SearchProductResponse {
	var res []SearchProductResponse
	for _, d := range *products {
		mapped := SearchProductResponse{
			ID:        d.ID,
			Name:      d.Name,
			Stock:     d.Stock,
			Images:    d.Images,
			Price:     d.Price,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
			User: userSearchProductResponse{
				ID:   d.User.ID,
				Name: d.User.Name,
			},
		}

		res = append(res, mapped)
	}

	return res
}
