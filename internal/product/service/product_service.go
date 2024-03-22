package service

import (
	"github.com/dafailyasa/learn-golang-template/internal/product/model"
)

type ProductService interface {
	Create(body *model.ProductCreateRequest, userEmail string) error
	Search(email string, params *model.ProductSearchParams) ([]model.SearchProductResponse, int64, error)
}
