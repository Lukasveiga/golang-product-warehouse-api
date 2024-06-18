package usecase

import (
	"fmt"
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
	"product-warehouse/internal/repository/inMemory"
	"product-warehouse/internal/usecase/dto"
	"testing"
)

func setup() (port.StockRepository, port.ProductRepository, *CreateStockUsecase) {
	stockRepository := inMemory.NewInMemoryStockRepository()
	productRepository := inMemory.NewInMemoryProductRepository()
	createStockUsecase := NewCreateStockUsecase(stockRepository, productRepository)

	return stockRepository, productRepository, createStockUsecase
} 

func TestStockUsecase_Success_CreateStock(t *testing.T) {
	_, productRepository, createStockUsecase := setup()

	productDto := domain.Product{
		Name: "Product Dto Test",
		Description: "Product Dto Test Description",
		Price: 5.5,
	}

	product := productRepository.AddProduct(&productDto)

	stockDtoTest := dto.StockDto{
		Product_id: product.Id,
		Quantity: 10,
	}

	resultTest, err := createStockUsecase.Execute(&stockDtoTest)

	if err != nil {
		t.Errorf("Error expected: %v, got: %v", nil, err)
	}

	if !stockEquity(dto.NewStock(&stockDtoTest), resultTest) {
		t.Errorf("Stock body expected: %v, got: %v", stockDtoTest, resultTest)
	}

	if resultTest.Id == 0 {
		t.Errorf("Stock id expected to be different from 0")
	}
}

func TestStockUsecase_SuccessExistingStock_CreateStock(t *testing.T) {
	stockRepository, productRepository, createStockUsecase := setup()

	productDto := domain.Product{
		Name: "Product Dto Test",
		Description: "Product Dto Test Description",
		Price: 5.5,
	}

	product := productRepository.AddProduct(&productDto)

	stockTest := domain.Stock{
		Product_id: product.Id,
		Quantity: 10,
	}

	stockRepository.AddStock(&stockTest)

	stockDtoTest := dto.StockDto{
		Product_id: product.Id,
		Quantity: 10,
	}

	resultTest, err := createStockUsecase.Execute(&stockDtoTest)

	if err != nil {
		t.Errorf("Error expected: %v, got: %v", nil, err)
	}

	if resultTest.Quantity != 20 {
		t.Errorf("Quantity expected: %d, got: %v", 20, resultTest.Quantity)
	}
}

func TestStockUsecase_ProductNotFound_CreateStock(t *testing.T) {
	_, _, createStockUsecase := setup()

	stockDtoTest := dto.StockDto{
		Product_id: 2,
		Quantity: 10,
	}

	expectedError := fmt.Errorf("product with id %d not found", 2)

	resultTest, err := createStockUsecase.Execute(&stockDtoTest)

	if resultTest != nil {
		t.Errorf("Result test expected: %v, got: %v", nil, resultTest)
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, got: %v", expectedError, err)
	}
}

func stockEquity(s1 *domain.Stock, s2 *domain.Stock) bool {
	if s1.Product_id != s2.Product_id {
		return false
	}

	if s1.Quantity != s2.Quantity {
		return false
	}

	return true
}