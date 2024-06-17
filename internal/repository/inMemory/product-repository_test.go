package inMemory

import (
	"fmt"
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
	"reflect"
	"testing"
)

var productRepository port.ProductRepository

func TestProductRepository_Success_AddProduct(t *testing.T) {
	productRepository = NewInMemoryProductRepository()

	productTest := domain.Product{
		Name: "Test Product Name",
		Description: "Test Product Description",
		Price: 2.5, 
	}

	resultTest := productRepository.AddProduct(&productTest)

	if !reflect.DeepEqual(*resultTest, productTest) {
		t.Errorf("Product body expected: %v, got: %v", resultTest, productTest)
	}
}

func TestProductRepository_Success_FindProductById(t *testing.T) {
	productRepository = NewInMemoryProductRepository()

	productTest := domain.Product{
		Name: "Test Product Name",
		Description: "Test Product Description",
		Price: 2.5, 
	}

	productRepository.AddProduct(&productTest)

	resultTest, _ := productRepository.FindProductById(productTest.Id)

	if !reflect.DeepEqual(*resultTest, productTest) {
		t.Errorf("Product body expected: %v, got: %v", resultTest, productTest)
	}
}

func TestProductRepository_NotFound_FindProductById(t *testing.T) {
	productRepository = NewInMemoryProductRepository()

	productId := 1

	errorTest := fmt.Errorf("product with id %d not found", productId)

	_, err := productRepository.FindProductById(productId)

	if err.Error() != errorTest.Error() {
		t.Errorf("Error message expected: %v, got: %v", errorTest, err)
	}
}