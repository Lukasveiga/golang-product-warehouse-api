package usecase

import (
	"fmt"
	"product-warehouse/internal/domain"
	"product-warehouse/internal/repository/inMemory"
	"reflect"
	"testing"
)

func TestProductUsecase_Success_FindProductById(t *testing.T) {
	productRepository := inMemory.NewInMemoryProductRepository()
	findProductByIdUsecase := NewFindProductByIdUsecase(productRepository)

	productTest := domain.Product{
		Name: "Test Product Name",
		Description: "Test Product Description",
		Price: 2.5,
	}

	productRepository.AddProduct(&productTest)

	resultTest, err := findProductByIdUsecase.Execute(productTest.Id)

	if err != nil {
		t.Errorf("Error expected to be nil")
	}

	if !reflect.DeepEqual(*resultTest, productTest) {
		t.Errorf("Product body expected: %v, got: %v", productTest, resultTest)
	}
}

func TestProductUsecase_NotFound_FindProductById(t *testing.T) {
	productRepository := inMemory.NewInMemoryProductRepository()
	findProductByIdUsecase := NewFindProductByIdUsecase(productRepository)

	productId := 1

	expectedError := fmt.Errorf("product with id %d not found", productId)

	resultTest, err := findProductByIdUsecase.Execute(productId)

	if resultTest != nil {
		t.Errorf("Result test expected to be: %v got: %v", nil, resultTest)
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("Error expected: %v, got: %v", expectedError, err)
	}
}