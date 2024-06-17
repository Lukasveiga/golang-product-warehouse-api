package usecase

import (
	"product-warehouse/internal/domain"
	"product-warehouse/internal/domain/dto"
	port "product-warehouse/internal/port/repository"
)

type CreateProductUsecase struct {
	repo port.ProductRepository
}

func NewCreateProductUsecase(repo port.ProductRepository) *CreateProductUsecase {
	return &CreateProductUsecase{
		repo: repo,
	}
}

func (uc CreateProductUsecase) Execute(product dto.ProductDto) *domain.Product {
	return uc.repo.AddProduct(dto.NewProduct(product))
}