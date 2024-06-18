package dto

import (
	"product-warehouse/internal/domain"
	"reflect"
	"testing"
)

func TestProductDto_Success_NewProduct(t *testing.T) {
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

	if !reflect.DeepEqual(*resultTest, productTest) {
		t.Errorf("Product body expected: %v, got: %v", productTest, resultTest)
	}
}