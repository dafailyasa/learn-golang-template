package service

import (
	"github.com/dafailyasa/learn-golang-template/internal/product/entity"
	"github.com/dafailyasa/learn-golang-template/internal/product/model"
)

type ProductService interface {
	Create(body *model.ProductCreateRequest, userEmail string) error
	Search(email string, params *model.ProductSearchParams) (*[]entity.Product, int64, error)
}
