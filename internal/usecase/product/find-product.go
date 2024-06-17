package usecase

import (
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
	product, err := uc.repo.FindProductById(productId)
	return product, err
}