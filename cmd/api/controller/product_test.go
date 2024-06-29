package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"product-warehouse/internal/domain"
	"product-warehouse/internal/usecase/dto"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)



func TestProductController(t *testing.T) {

	product := &domain.Product{
		Id: 1,
		Name: "Product Test",
		Description: "Product Description",
		Price: 20.0,
	}

	mockProduct := &MockProduct{Product: product}

	productController := ProductControllerSetup(mockProduct)

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

	t.Run("Create BadRequest Decoding Error", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/product", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()

		productController.Create(res, req) 

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
		assert.Equal(t, "Decoding Error\n", res.Body.String())
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

	t.Run("FindById Success", func(t *testing.T){
		req := httptest.NewRequest("GET", "/product", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		res := httptest.NewRecorder()

		productController.FindById(res, req)

		got := res.Result()

		var responseBody domain.Product
		json.NewDecoder(res.Body).Decode(&responseBody)

		assert.Equal(t, http.StatusOK, got.StatusCode)
		assert.Equal(t, *product, responseBody)
	})

	t.Run("FindById NotFound Product", func(t *testing.T) {
		mockProduct = &MockProduct{Product: nil}
		productController = ProductControllerSetup(mockProduct)

		id := "2"

		req := httptest.NewRequest("GET", "/product", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", id)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		res := httptest.NewRecorder()

		productController.FindById(res, req)

		assert.Equal(t, http.StatusNotFound, res.Result().StatusCode)
		assert.Equal(t, fmt.Sprintf("product not found with id %s\n", id), res.Body.String())
	})

	t.Run("FindById Invalid Id Param", func(t *testing.T) {
		mockProduct = &MockProduct{Product: nil}
		productController = ProductControllerSetup(mockProduct)

		id := "a"

		req := httptest.NewRequest("GET", "/product", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", id)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		res := httptest.NewRecorder()

		productController.FindById(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
		assert.Equal(t, "Invalid id\n", res.Body.String())
	})
}