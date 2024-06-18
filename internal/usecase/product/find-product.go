package usecase

import (
	"fmt"
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
)

type FindProductByIdUsecase struct {
	repo port.ProductRepository
}

func NewFindProductByIdUsecase(repo port.ProductRepository) *FindProductByIdUsecase {
	return &FindProductByIdUsecase{
		repo: repo,
	}
}

func (uc FindProductByIdUsecase) Execute(productId int) (*domain.Product, error) {
	product := uc.repo.FindProductById(productId)

	if product == nil {
		return nil, fmt.Errorf("product with id %d not found", productId)
	}
	
	return product, nil
}