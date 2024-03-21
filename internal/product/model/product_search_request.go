package model

type ProductSearchParams struct {
	Page   int    `json:"page"`
	Size   int    `json:"size"`
	Search string `json:"search"`
}

func (s *ProductSearchParams) GetOffset() int {
	return (s.Page - 1) * s.Size
}

func (s *ProductSearchParams) GetSize() int {
	return s.Size
}
