package usecase

import (
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
	"product-warehouse/internal/repository/inMemory"
	"product-warehouse/internal/shared"
	"product-warehouse/internal/usecase/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createProductUsecaseSetup() (port.ProductRepository, *CreateProductUsecase) {
	productRepository := inMemory.NewInMemoryProductRepository()
	createProductUsecase := NewCreateProductUsecase(productRepository)

	return productRepository, createProductUsecase
}

func TestCreateProductUsecase(t *testing.T) {

	_, createProductUsecase := createProductUsecaseSetup()

	t.Run("CreateProduct Success", func(t *testing.T) {
		productDtoTest := dto.ProductDto{
			Name: "Product Dto Test",
			Description: "Product Dto Test Description",
			Price: 5.5,
		}
	
		resultTest, err := createProductUsecase.Execute(&productDtoTest)

		product, _ := dto.NewProduct(&productDtoTest)
	
		assert.Nil(t, err)
		assert.True(t, productEquity(resultTest, product))
		assert.NotEqual(t, 0, resultTest.Id)
	})

	t.Run("CreateProduct Invalid Inputs", func(t *testing.T) {
		productDtoTest := dto.ProductDto{
			Name: "",
			Description: "",
			Price: -1.0,
		}

		expectedError := &shared.ValidationError{
			Errors: map[string]string{
				"name": "cannot be empty",
				"description": "cannot be empty",
				"price": "must be greater than zero",
			},
		}

		resultTest, err := createProductUsecase.Execute(&productDtoTest)

		_, ok := err.(*shared.ValidationError)

		assert.Nil(t, resultTest)
		assert.True(t, ok)
		assert.Equal(t, expectedError, err)
	})
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