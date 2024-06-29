package dto

import (
	"product-warehouse/internal/domain"
	"product-warehouse/internal/shared"
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
	
		resultTest, err := NewProduct(&productDtoTest)
	
		assert.Nil(t, err)
		assert.Equal(t, productTest, *resultTest)
	})

	t.Run("NewProduct Invalid Input", func(t *testing.T) {
		productDtoTest := ProductDto{
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

		resultTest, err := NewProduct(&productDtoTest)

		_, ok := err.(*shared.ValidationError)

		assert.Nil(t, resultTest)
		assert.True(t, ok)
		assert.Equal(t, expectedError, err)
	})
}