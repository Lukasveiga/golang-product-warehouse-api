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

func TestStockController(t *testing.T) {

	product := &domain.Product{
		Id: 1,
		Name: "Product Test",
		Description: "Product Description",
		Price: 20.0,
	}

	stock := &domain.Stock{
		Id: 1,
		Product_id: product.Id,
		Quantity: 10,
	}

	mockProduct := &MockProduct{Product: product}
	mockStock := &MockStock{Stock: stock}

	stockController := StockControllerSetup(mockStock, mockProduct)

	t.Run("Create Success", func(t *testing.T) {

		stockDto := &dto.StockDto{
			Product_id: product.Id,
			Quantity: 10,
		}

		body, _ := json.Marshal(stockDto)

		req := httptest.NewRequest("POST", "/stock", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		stockController.Create(res, req)

		got := res.Result()

		var responseBody domain.Stock
		json.NewDecoder(res.Body).Decode(&responseBody)

		assert.Equal(t, http.StatusCreated, got.StatusCode)
		assert.Equal(t, *stock, responseBody)
	})

	t.Run("Create BadRequest Decoding Error", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/stock", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		res := httptest.NewRecorder()

		stockController.Create(res, req) 

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
		assert.Equal(t, "Decoding Error\n", res.Body.String())
	})

	t.Run("Create BadRequest Invalid Input", func(t *testing.T) {

		stockDto := &dto.StockDto{
			Product_id: product.Id,
			Quantity: -1,
		}

		body, _ := json.Marshal(stockDto)

		req := httptest.NewRequest("POST", "/stock", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		stockController.Create(res, req)

		got := res.Result()

		var responseBody map[string]string
		json.NewDecoder(res.Body).Decode(&responseBody)

		assert.Equal(t, http.StatusBadRequest, got.StatusCode)
		assert.Equal(t, responseBody["quantity"], "cannot be negative value")
	})

	t.Run("Create Product NotFound", func(t *testing.T) {
		mockProduct = &MockProduct{Product: nil}

		stockController := StockControllerSetup(mockStock, mockProduct)

		productId := 1

		stockDto := &dto.StockDto{
			Product_id: productId,
			Quantity: 1,
		}

		body, _ := json.Marshal(stockDto)

		req := httptest.NewRequest("POST", "/stock", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		stockController.Create(res, req)

		got := res.Result()

		assert.Equal(t, http.StatusNotFound, got.StatusCode)
		assert.Equal(t, fmt.Sprintf("product not found with id %d\n", productId), res.Body.String())
	})

	t.Run("FindStockByProductId Success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/stock", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("productId", "1")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		res := httptest.NewRecorder()

		stockController.FindStockByProductId(res, req)

		got := res.Result()

		var responseBody domain.Stock
		json.NewDecoder(res.Body).Decode(&responseBody)

		assert.Equal(t, http.StatusOK, got.StatusCode)
		assert.Equal(t, *stock, responseBody)
	})

	t.Run("FindStockByProductId NotFound Stock", func(t *testing.T) {
		mockStock = &MockStock{Stock: nil}

		stockController := StockControllerSetup(mockStock, mockProduct)

		productId := "1"

		req := httptest.NewRequest("GET", "/stock", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("productId", productId)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		res := httptest.NewRecorder()

		stockController.FindStockByProductId(res, req)

		got := res.Result()

		assert.Equal(t, http.StatusNotFound, got.StatusCode)
		assert.Equal(t, fmt.Sprintf("product not found with id %s\n", productId), res.Body.String())
	})

	t.Run("FindStockByProductId Invalid Id Param", func(t *testing.T) {
		productId := "a"

		req := httptest.NewRequest("GET", "/product", nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", productId)
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		res := httptest.NewRecorder()

		stockController.FindStockByProductId(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Result().StatusCode)
		assert.Equal(t, "Invalid productId\n", res.Body.String())
	})
}