package usecase

import (
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
	"product-warehouse/internal/shared"
	"product-warehouse/internal/usecase/dto"
)

type CreateProductUsecase struct {
	repo port.ProductRepository
}

func NewCreateProductUsecase(repo port.ProductRepository) *CreateProductUsecase {
	return &CreateProductUsecase{
		repo: repo,
	}
}

func (uc CreateProductUsecase) Execute(product *dto.ProductDto) (*domain.Product, shared.ErrorMap) {
	p, err := dto.NewProduct(product)

	if err != nil {
		return nil, err
	}
	
	return uc.repo.AddProduct(p), nil
}