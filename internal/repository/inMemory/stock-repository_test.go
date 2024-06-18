package inMemory

import (
	"fmt"
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
	"reflect"
	"testing"
)

var stockRepository port.StockRepository

func TestStockRepository_Success_AddStock(t *testing.T) {
	stockRepository = NewInMemoryStockRepository()

	stockTest := domain.Stock{
		Product_id: 1,
		Quantity: 20,
	}

	resultTest := stockRepository.AddStock(&stockTest);

	if !reflect.DeepEqual(*resultTest, stockTest) {
		t.Errorf("Stock body expected: %v, got: %v", stockTest, resultTest)
	}
}

func TestStockRepository_Success_FindStockByProductId(t *testing.T) {
	stockRepository = NewInMemoryStockRepository()

	stockTest := domain.Stock{
		Product_id: 1,
		Quantity: 20,
	}

	stockRepository.AddStock(&stockTest)

	resultTest, _ := stockRepository.FindStockByProductId(stockTest.Product_id)

	if !reflect.DeepEqual(*resultTest, stockTest) {
		t.Errorf("Stock body expected: %v, got: %v", stockTest, resultTest)
	}
}

func TestStockRepository_NotFound_FindStockByProductId(t *testing.T) {
	stockRepository = NewInMemoryStockRepository()

	productId := 1

	errorTest := fmt.Errorf("stock with productId %d not found", productId)

	_, resultTest := stockRepository.FindStockByProductId(productId)

	if resultTest.Error() != errorTest.Error() {
		t.Errorf("Error message expected: %v, got: %v", errorTest,resultTest)
	}
}

func TestStockRepository_Success_UpdateStockQuantity(t *testing.T) {
	stockRepository = NewInMemoryStockRepository()

	stockTest := domain.Stock{
		Product_id: 1,
		Quantity: 20,
	}

	newQuantity := 10

	stockRepository.AddStock(&stockTest)

	resultTest := stockRepository.UpdateStockQuantity(stockTest.Id, newQuantity)

	if resultTest.Quantity != 30 {
		t.Errorf("Quantity value expected: %d, got: %d", stockTest.Quantity + newQuantity, resultTest.Quantity)
	}
}