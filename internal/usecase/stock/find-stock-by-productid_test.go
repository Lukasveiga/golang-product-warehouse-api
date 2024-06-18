package usecase

import (
	"fmt"
	"product-warehouse/internal/domain"
	"product-warehouse/internal/repository/inMemory"
	"reflect"
	"testing"
)

func TestStockUsecase_Success_FindStockByProductId(t *testing.T) {
	stockRepository := inMemory.NewInMemoryStockRepository()
	findStockByProductIdUsecase := NewFindStockByProductIdUsecase(stockRepository)

	stockTest := domain.Stock{
		Product_id: 1,
		Quantity: 10,
	}

	stockRepository.AddStock(&stockTest)

	resultTest, err := findStockByProductIdUsecase.Execute(stockTest.Id)

	if err != nil {
		t.Errorf("Error expected to be nil")
	}

	if !reflect.DeepEqual(stockTest, *resultTest) {
		t.Errorf("Stock body expected: %v, got: %v", stockTest, resultTest)
	}
}

func TestStockUsecase_NotFound_FindStockByProductId(t *testing.T) {
	stockRepository := inMemory.NewInMemoryStockRepository()
	findStockByProductIdUsecase := NewFindStockByProductIdUsecase(stockRepository)

	productId := 2

	expectedError := fmt.Errorf("product with id %d not found", productId)

	resultTest, err := findStockByProductIdUsecase.Execute(productId)

	if resultTest != nil {
		t.Errorf("Result expected: %v, got: %v", nil, resultTest)
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("Error expected: %v, got: %v", expectedError, err)
	}
}