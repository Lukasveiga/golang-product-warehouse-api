package inMemory

import (
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func stockRepositorySetup() (port.StockRepository, *domain.Stock) {
	stockTest := domain.Stock{
		Product_id: 1,
		Quantity: 20,
	}

	return NewInMemoryStockRepository(), &stockTest
}

func TestStockRepository(t *testing.T) {

	stockRepository, stockTest := stockRepositorySetup()

	t.Run("AddStock Success", func(t *testing.T) {
		resultTest := stockRepository.AddStock(stockTest)

		assert.Equal(t, stockTest, resultTest)
	})

	t.Run("FindStockByProductId Success", func(t *testing.T) {
		resultTest := stockRepository.FindStockByProductId(stockTest.Product_id)

		assert.Equal(t, stockTest, resultTest)
	})

	t.Run("FindStockByProductId Not Found", func(t *testing.T) {
		invalidProductId := stockTest.Product_id + 1

		resultTest := stockRepository.FindStockByProductId(invalidProductId)

		assert.Nil(t, resultTest)
	})

	t.Run("UpdateStockQuantity Success", func(t *testing.T) {
		newQuantity := 10

		stockRepository.AddStock(stockTest)

		resultTest := stockRepository.UpdateStockQuantity(stockTest.Id, newQuantity)

		assert.Equal(t, 10, resultTest.Quantity)
	})
}