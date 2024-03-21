package repository

import (
	"github.com/dafailyasa/learn-golang-template/internal/product/entity"
	"github.com/dafailyasa/learn-golang-template/internal/product/model"
)

type ProductRepository interface {
	FindProductByUserId(userId string) (*[]entity.Product, error)
	Create(*entity.Product) error
	FindById(id string) (*entity.Product, error)
	Search(userId string, params *model.ProductSearchParams) (*[]entity.Product, int64, error)
}
