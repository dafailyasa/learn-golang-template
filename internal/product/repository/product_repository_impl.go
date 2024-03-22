package repository

import (
	"strings"

	"github.com/dafailyasa/learn-golang-template/internal/product/entity"
	"github.com/dafailyasa/learn-golang-template/internal/product/model"
	"gorm.io/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{
		DB: db,
	}
}

var _ ProductRepository = (*productRepository)(nil)

func (p *productRepository) FindProductByUserId(userId string) (*[]entity.Product, error) {
	var product []entity.Product

	if err := p.DB.Model(&entity.Product{}).Where("user_id = ?", userId).Find(&product).Error; err != nil {
		return &product, err
	}

	return &product, nil
}

func (p *productRepository) Create(product *entity.Product) error {
	if err := p.DB.Model(&entity.Product{}).Create(product).Error; err != nil {
		return err
	}

	return nil
}

func (p *productRepository) FindById(id string) (*entity.Product, error) {
	var product entity.Product

	if err := p.DB.Model(&entity.Product{}).Where("id = ?", id).Find(&product).Error; err != nil {
		return &product, err
	}

	return &product, nil
}

func (p *productRepository) Search(userId string, params *model.ProductSearchParams) (*[]entity.Product, int64, error) {
	var products *[]entity.Product
	var total int64 = 0

	filter := p.filterSearchQuery(userId, params)

	if err := p.DB.Model(&entity.Product{}).Scopes(filter).Preload("User").Offset(params.GetOffset()).Limit(params.GetSize()).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	if err := p.DB.Model(&entity.Product{}).Scopes(filter).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (p *productRepository) filterSearchQuery(userId string, params *model.ProductSearchParams) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("user_id = ?", userId)
		if search := strings.TrimSpace(params.Search); search != "" {
			search = "%" + search + "%"
			tx.Where("name LIKE ?", search)
		}

		return tx
	}
}
