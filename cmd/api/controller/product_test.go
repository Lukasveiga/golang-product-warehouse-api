package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"product-warehouse/internal/domain"
	"product-warehouse/internal/usecase/dto"
	usecase "product-warehouse/internal/usecase/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockProduct struct {
	product *domain.Product
}

func(m *MockProduct) AddProduct(product *domain.Product) *domain.Product {
	return m.product
}

func (m *MockProduct) FindProductById(id int) (*domain.Product) {
	return m.product
}

func productControllerSetup(mock *MockProduct) *ProductController{
	createProductUsecase := usecase.NewCreateProductUsecase(mock)
	findProductUsecase := usecase.NewFindProductByIdUsecase(mock)
	return NewProductController(createProductUsecase, findProductUsecase)
}

func TestProductController(t *testing.T) {

	product := &domain.Product{
		Id: 1,
		Name: "Product Test",
		Description: "Product Description",
		Price: 20.0,
	}

	mockProduct := &MockProduct{product: product}

	productController := productControllerSetup(mockProduct)

	t.Run("Create Success", func(t *testing.T) {

		productDto := &dto.ProductDto{
			Name: product.Name,
			Description: product.Description,
			Price: product.Price,
		}

		body, _  := json.Marshal(productDto)

		req := httptest.NewRequest("POST", "/product", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		productController.Create(res, req)

		got := res.Result()

		var responseBody domain.Product
		json.NewDecoder(res.Body).Decode(&responseBody)

		assert.Equal(t, http.StatusCreated, got.StatusCode)
		assert.Equal(t, *product, responseBody)
	})

	t.Run("Create BadRequest Invalid Input", func(t *testing.T) {

		productDto := &dto.ProductDto{
			Name: "",
			Description: "",
			Price: -1.0,
		}

		body, _ := json.Marshal(productDto)

		req := httptest.NewRequest("POST", "/product", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		productController.Create(res, req)

		got := res.Result()

		var responseBody map[string]string
		json.NewDecoder(res.Body).Decode(&responseBody)

		assert.Equal(t, http.StatusBadRequest, got.StatusCode)
		assert.Equal(t, responseBody["name"], "cannot be empty")
		assert.Equal(t, responseBody["description"], "cannot be empty")
		assert.Equal(t, responseBody["price"], "must be greater than zero")
	})
}