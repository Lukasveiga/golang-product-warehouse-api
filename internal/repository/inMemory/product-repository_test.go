package inMemory

import (
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func productRepositorySetup() (port.ProductRepository, *domain.Product) {
	productTest := domain.Product{
		Name: "Test Product Name",
		Description: "Test Product Description",
		Price: 2.5, 
	}

	return NewInMemoryProductRepository(), &productTest
}

func TestProductRepository(t *testing.T) {

	productRepository, productTest := productRepositorySetup()

	t.Run("AddProduct Success", func(t *testing.T) {
		resultTest := productRepository.AddProduct(productTest)

		assert.Equal(t, productTest, resultTest)
	})

	t.Run("FindProductById Success", func(t *testing.T) {
		productRepository.AddProduct(productTest)

		resultTest := productRepository.FindProductById(productTest.Id)

		assert.Equal(t, productTest, resultTest)
	})

	t.Run("FindProductById Not Found", func(t *testing.T) {
		invalidProductId := productTest.Id + 1

		resultTest := productRepository.FindProductById(invalidProductId)

		assert.Nil(t, resultTest)
	})
}