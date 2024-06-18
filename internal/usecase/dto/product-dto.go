package dto

import "product-warehouse/internal/domain"

type ProductDto struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}

func NewProduct(productDto *ProductDto) *domain.Product {
	return &domain.Product{
		Name: productDto.Name,
		Description: productDto.Description,
		Price: productDto.Price,
	}
}