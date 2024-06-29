package usecase

import (
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
	"product-warehouse/internal/shared"
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
		return nil, &shared.NotFoundError{
			Object: "product",
			Id: productId,
		} 
	}
	return product, nil
}