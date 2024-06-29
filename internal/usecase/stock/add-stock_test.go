package usecase

import (
	"fmt"
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
	"product-warehouse/internal/repository/inMemory"
	"product-warehouse/internal/shared"
	"product-warehouse/internal/usecase/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func addStockUsecaseSetup() (port.StockRepository, port.ProductRepository, *CreateStockUsecase, *domain.Product) {
	stockRepository := inMemory.NewInMemoryStockRepository()
	productRepository := inMemory.NewInMemoryProductRepository()	
	createStockUsecase := NewCreateStockUsecase(stockRepository, productRepository)

	productTest := domain.Product{
		Name: "Product Dto Test",
		Description: "Product Dto Test Description",
		Price: 5.5,
	}

	return stockRepository, productRepository, createStockUsecase, &productTest
} 

func TestAddStockUsecase(t *testing.T) {

	stockRepository, productRepository, createStockUsecase, productTest := addStockUsecaseSetup()

	t.Run("CreateStock Success", func(t *testing.T) {
		product := productRepository.AddProduct(productTest)

		stockDtoTest := dto.StockDto{
			Product_id: product.Id,
			Quantity: 10,
		}

		resultTest, err := createStockUsecase.Execute(&stockDtoTest)

		stock, _ := dto.NewStock(&stockDtoTest)

		assert.Nil(t, err)
		assert.NotEqual(t, 0, resultTest.Id)
		assert.True(t, stockEquity(stock, resultTest))
	})

	t.Run("CreateStock Success Existing Stock", func(t *testing.T) {
		product := productRepository.AddProduct(productTest)

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

		assert.Nil(t, err)
		assert.Equal(t, 20, resultTest.Quantity)
	})

	t.Run("CreateStock Product Not Found", func(t *testing.T) {
		stockDtoTest := dto.StockDto{
			Product_id: 3,
			Quantity: 10,
		}

		expectedError := fmt.Errorf("product not found with id %d", stockDtoTest.Product_id)

		resultTest, err := createStockUsecase.Execute(&stockDtoTest)

		_, ok := err.(*shared.NotFoundError)

		assert.Nil(t, resultTest)
		assert.True(t, ok)
		assert.Equal(t, expectedError.Error(), err.Error())
	})

	t.Run("CreateStock Invalid Input", func(t *testing.T) {
		product := productRepository.AddProduct(productTest)

		stockDtoTest := dto.StockDto{
			Product_id: product.Id,
			Quantity: -1,
		}

		expectedError := &shared.ValidationError{
			Errors: map[string]string{
				"quantity": "cannot be negative value",
			},
		}

		resultTest, err := createStockUsecase.Execute(&stockDtoTest)

		_, ok := err.(*shared.ValidationError)

		assert.Nil(t, resultTest)
		assert.True(t, ok)
		assert.Equal(t, expectedError, err)
	})
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