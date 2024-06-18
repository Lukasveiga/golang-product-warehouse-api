package inMemory

import (
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
		t.Errorf("Product body expected: %v, got: %v", productTest, resultTest)
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

	resultTest := productRepository.FindProductById(productTest.Id)

	if !reflect.DeepEqual(*resultTest, productTest) {
		t.Errorf("Product body expected: %v, got: %v", productTest, resultTest)
	}
}

func TestProductRepository_NotFound_FindProductById(t *testing.T) {
	productRepository = NewInMemoryProductRepository()

	productId := 1

	resultTest := productRepository.FindProductById(productId)

	if resultTest != nil {
		t.Errorf("Error message expected: %v, got: %v", nil, resultTest)
	}
}