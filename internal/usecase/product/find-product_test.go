package usecase

import (
	"fmt"
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
	"product-warehouse/internal/repository/inMemory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func productUsecaseSetup() (port.ProductRepository, *FindProductByIdUsecase, *domain.Product) {
	productRepository := inMemory.NewInMemoryProductRepository()
	findProductByIdUsecase := NewFindProductByIdUsecase(productRepository)

	productTest := domain.Product{
		Name: "Test Product Name",
		Description: "Test Product Description",
		Price: 2.5,
	}

	return productRepository, findProductByIdUsecase, &productTest
}

func TestProductUsecase_Success_FindProductById(t *testing.T) {

	productRepository, findProductByIdUsecase, productTest := productUsecaseSetup()

	t.Run("FindProductById Success", func(t *testing.T) {
		productRepository.AddProduct(productTest)

		resultTest, err := findProductByIdUsecase.Execute(productTest.Id)

		assert.Nil(t, err)
		assert.Equal(t, *productTest, *resultTest)
	})

	t.Run("FindProductById Not Found", func(t *testing.T) {
		invalidProductId := productTest.Id + 1

		expectedError := fmt.Errorf("product with id %d not found", invalidProductId)

		resultTest, err := findProductByIdUsecase.Execute(invalidProductId)

		assert.Nil(t, resultTest)
		assert.Equal(t, expectedError, err)
	})
}