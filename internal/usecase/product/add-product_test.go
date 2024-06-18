package usecase

import (
	"product-warehouse/internal/domain"
	"product-warehouse/internal/repository/inMemory"
	"product-warehouse/internal/usecase/dto"
	"testing"
)

func TestProductUsecase_Success_CreateProduct(t *testing.T) {
	productRepository := inMemory.NewInMemoryProductRepository()
	createProductUsecase := NewCreateProductUsecase(productRepository)

	productDtoTest := dto.ProductDto{
		Name: "Product Dto Test",
		Description: "Product Dto Test Description",
		Price: 5.5,
	}

	resultTest := createProductUsecase.Execute(&productDtoTest)

	if !productEquity(resultTest, dto.NewProduct(&productDtoTest)) {
		t.Errorf("Product body expected: %v, got: %v", productDtoTest, resultTest)
	}

	if resultTest.Id == 0 {
		t.Errorf("Product id expected to be different from 0")
	}
}

func productEquity(p1 *domain.Product, p2 *domain.Product) bool {
	if p1.Name != p2.Name {
		return false
	}

	if p1.Description != p2.Description {
		return false
	}

	if p1.Price != p2.Price {
		return false
	}

	return true
}