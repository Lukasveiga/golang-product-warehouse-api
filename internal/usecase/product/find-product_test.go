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

func findProductUsecaseSetup() (port.ProductRepository, *FindProductByIdUsecase, *domain.Product) {
	productRepository := inMemory.NewInMemoryProductRepository()
	findProductByIdUsecase := NewFindProductByIdUsecase(productRepository)

	productTest := domain.Product{
		Name: "Test Product Name",
		Description: "Test Product Description",
		Price: 2.5,
	}

	return productRepository, findProductByIdUsecase, &productTest
}

func TestFindProductUsecase(t *testing.T) {

	productRepository, findProductByIdUsecase, productTest := findProductUsecaseSetup()

	t.Run("FindProductById Success", func(t *testing.T) {
		productRepository.AddProduct(productTest)

		resultTest, err := findProductByIdUsecase.Execute(productTest.Id)

		assert.Nil(t, err)
		assert.Equal(t, *productTest, *resultTest)
	})

	t.Run("FindProductById Not Found", func(t *testing.T) {
		invalidProductId := productTest.Id + 1

		expectedError := fmt.Errorf("product not found with id %d", invalidProductId)

		resultTest, err := findProductByIdUsecase.Execute(invalidProductId)

		_, ok := err.(*shared.NotFoundError)

		assert.Nil(t, resultTest)
		assert.True(t, ok)
		assert.Equal(t, expectedError.Error(), err.Error())
	})
}