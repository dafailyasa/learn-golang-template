package model

type ProductCreateRequest struct {
	Name        string   `json:"name" validate:"required"`
	Stock       uint32   `json:"stock" validate:"required,number,gt=0"`
	Price       float64  `json:"price" validate:"required,number,gt=0"`
	Images      []string `json:"images" validate:"required,min=1,dive"`
	IsPublished bool     `json:"isPublished"`
}
