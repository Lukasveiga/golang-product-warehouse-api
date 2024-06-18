package usecase

import (
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
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

func (uc CreateProductUsecase) Execute(product *dto.ProductDto) *domain.Product {
	return uc.repo.AddProduct(dto.NewProduct(product))
}