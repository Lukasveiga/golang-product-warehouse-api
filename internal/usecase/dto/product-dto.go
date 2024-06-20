package dto

import (
	"product-warehouse/internal/domain"
	"product-warehouse/internal/shared"
)

type ProductDto struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}

func NewProduct(productDto *ProductDto) (*domain.Product, shared.ErrorMap) {
	p := &domain.Product{
		Name: productDto.Name,
		Description: productDto.Description,
		Price: productDto.Price,
	}

	errs := p.Validate()

	if errs != nil {
		return nil, errs
	}

	return p, nil 
}