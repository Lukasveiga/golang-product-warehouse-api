package dto

import (
	"product-warehouse/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductDto(t *testing.T) {
	t.Run("NewProduct Success", func(t *testing.T) {
		productTest := domain.Product{
			Name: "Test Product Name",
			Description: "Test Product Description",
			Price: 2.5, 
		}
	
		productDtoTest := ProductDto{
			productTest.Name,
			productTest.Description,
			productTest.Price,
		}
	
		resultTest := NewProduct(&productDtoTest)
	
		assert.Equal(t, productTest, *resultTest)
	})
}