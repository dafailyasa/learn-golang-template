package service

import (
	"encoding/json"

	auth "github.com/dafailyasa/learn-golang-template/internal/auth/repository"
	"github.com/dafailyasa/learn-golang-template/internal/product/entity"
	"github.com/dafailyasa/learn-golang-template/internal/product/model"
	product "github.com/dafailyasa/learn-golang-template/internal/product/repository"
	customErr "github.com/dafailyasa/learn-golang-template/pkg/custom-errors"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type productService struct {
	DB          *gorm.DB
	ProductRepo product.ProductRepository
	AuthRepo    auth.AuthRepository
}

func NewProductService(productRepo product.ProductRepository, authRepo auth.AuthRepository, db *gorm.DB) *productService {
	return &productService{
		DB:          db,
		ProductRepo: productRepo,
		AuthRepo:    authRepo,
	}
}

func (s *productService) Create(body *model.ProductCreateRequest, email string) error {
	user, err := s.AuthRepo.FindOneByEmail(email)

	if err != nil {
		return err
	}

	if user == nil {
		return customErr.ErrUserNotFound
	}

	imagesJson, err := json.Marshal(body.Images)
	if err != nil {
		return err
	}

	product := new(entity.Product)
	product.Name = body.Name
	product.Price = body.Price
	product.Stock = body.Stock
	product.Images = datatypes.JSON(imagesJson)
	product.IsPublished = body.IsPublished
	product.UserID = user.ID

	if err := s.ProductRepo.Create(product); err != nil {
		return err
	}

	return nil
}

func (s *productService) Search(email string, params *model.ProductSearchParams) (*[]entity.Product, int64, error) {
	user, err := s.AuthRepo.FindOneByEmail(email)
	if err != nil {
		return nil, 0, err
	}

	if user == nil {
		return nil, 0, customErr.ErrUserNotFound
	}
	products, total, err := s.ProductRepo.Search(user.ID.String(), params)
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}
