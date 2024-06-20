package dto

import (
	"errors"
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

		expectedError := shared.ErrorMap{
			"name": errors.New("cannot be empty"),
			"description": errors.New("cannot be empty"),
			"price": errors.New("must be greater than zero"),
		}

		resultTest, err := NewProduct(&productDtoTest)

		assert.Nil(t, resultTest)
		assert.Equal(t, expectedError, err)
	})
}