package usecase

import (
	"product-warehouse/internal/domain"
	"product-warehouse/internal/repository/inMemory"
	"product-warehouse/internal/usecase/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProductUsecase(t *testing.T) {
	t.Run("CreateProduct Success", func(t *testing.T) {
		productRepository := inMemory.NewInMemoryProductRepository()
		createProductUsecase := NewCreateProductUsecase(productRepository)
	
		productDtoTest := dto.ProductDto{
			Name: "Product Dto Test",
			Description: "Product Dto Test Description",
			Price: 5.5,
		}
	
		resultTest := createProductUsecase.Execute(&productDtoTest)
	
		assert.True(t, productEquity(resultTest, dto.NewProduct(&productDtoTest)))
	
		assert.NotEqual(t, 0, resultTest.Id)
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