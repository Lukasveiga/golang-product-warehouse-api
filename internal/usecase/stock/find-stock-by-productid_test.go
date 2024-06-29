package usecase

import (
	"fmt"
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
	"product-warehouse/internal/repository/inMemory"
	"product-warehouse/internal/shared"
	"testing"

	"github.com/stretchr/testify/assert"
)

func findStockByProductIdUsecaseSetup() (port.StockRepository, *FindStockByProductIdUsecase) {
	stockRepository := inMemory.NewInMemoryStockRepository()
	findStockByProductIdUsecase := NewFindStockByProductIdUsecase(stockRepository)

	return stockRepository, findStockByProductIdUsecase
}

func TestFindStockUsecase(t *testing.T) {

	stockRepository, findStockByProductIdUsecase := findStockByProductIdUsecaseSetup()

	t.Run("FindStockByProductId Success", func(t *testing.T) {
		stockTest := domain.Stock{
			Product_id: 1,
			Quantity: 10,
		}

		stockRepository.AddStock(&stockTest)

		resultTest, err := findStockByProductIdUsecase.Execute(stockTest.Id)

		assert.Nil(t, err)
		assert.Equal(t, stockTest, *resultTest)
	})

	t.Run("FindStockByProductId Not Found", func(t *testing.T) {
		invalidProductId := 2

		expectedError := fmt.Errorf("product not found with id %d", invalidProductId)

		resultTest, err := findStockByProductIdUsecase.Execute(invalidProductId)

		_, ok := err.(*shared.NotFoundError)

		assert.Nil(t, resultTest)
		assert.True(t, ok)
		assert.Equal(t, expectedError.Error(), err.Error())
	})
}