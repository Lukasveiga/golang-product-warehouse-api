package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"product-warehouse/config"
	"product-warehouse/internal/domain"
	"product-warehouse/internal/repository/postgre"
	"product-warehouse/internal/usecase/dto"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIntegrat(t *testing.T) {
	
	dbConnection := config.SetupDbConnection()

	productRepository := postgre.NewPostgreProductRepository(dbConnection)
	stockRepository := postgre.NewPostgreStockRepository(dbConnection)

	defer dbConnection.Close()

	go startServer("8080", "test", productRepository, stockRepository)

	url := "http://localhost:8080/api/v1"

	productDto := &dto.ProductDto{
		Name: "Product Test",
		Description: "Product Description",
		Price: 20.0,
	}

	stockDto := &dto.StockDto{
		Product_id: 1,
		Quantity: 10,
	}

	time.Sleep(2 * time.Second)

	t.Run("Should return status code 200 and product body", func(t *testing.T) {
		body, _  := json.Marshal(productDto)

		req, _ := http.NewRequest("POST", url + "/product", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var responseProduct domain.Product

		err = json.Unmarshal(responseBody, &responseProduct)

		if err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.Equal(t, productDto.Name, responseProduct.Name)
		assert.Equal(t, productDto.Description, responseProduct.Description)
		assert.Equal(t, productDto.Price, responseProduct.Price)
	})

	t.Run("Should return status code 400 and erros body", func(t *testing.T) {
		invalidProductDto := &dto.ProductDto{
			Name: "",
			Description: "",
			Price: -1.0,
		}

		body, _  := json.Marshal(invalidProductDto)

		req, _ := http.NewRequest("POST", url + "/product", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var responseErros map[string]string

		err = json.Unmarshal(responseBody, &responseErros)

		if err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Equal(t, responseErros["name"], "cannot be empty")
		assert.Equal(t, responseErros["description"], "cannot be empty")
		assert.Equal(t, responseErros["price"], "must be greater than zero")
	})

	t.Run("Should return status code 200 and product body", func(t *testing.T) {
		productId := "1"

		req, _ := http.NewRequest("GET", url + "/product/" + productId, nil)

		client := &http.Client{}
		res, err := client.Do(req)
		
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var responseProduct domain.Product

		err = json.Unmarshal(responseBody, &responseProduct)

		if err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, productDto.Name, responseProduct.Name)
		assert.Equal(t, productDto.Description, responseProduct.Description)
		assert.Equal(t, productDto.Price, responseProduct.Price)
	})

	t.Run("Should return status code 404 and erro body", func(t *testing.T) {
		productId := "2"

		req, _ := http.NewRequest("GET", url + "/product/" + productId, nil)

		client := &http.Client{}
		res, err := client.Do(req)
		
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
		assert.Equal(t, fmt.Sprintf("product with id %s not found\n", productId), string(responseBody))
	})

	t.Run("Should return status code 201 and stock body", func(t *testing.T) {
		body, _  := json.Marshal(stockDto)

		req, _ := http.NewRequest("POST", url + "/stock", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var responseStock domain.Stock

		err = json.Unmarshal(responseBody, &responseStock)

		if err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusCreated, res.StatusCode)
		assert.Equal(t, stockDto.Product_id, responseStock.Product_id)
		assert.Equal(t, stockDto.Quantity, responseStock.Quantity)
	})

	t.Run("Should return status code 400 and erros body", func(t *testing.T) {
		invalidStockDto := &dto.StockDto{
			Product_id: 1,
			Quantity: -1,
		}

		body, _  := json.Marshal(invalidStockDto)

		req, _ := http.NewRequest("POST", url + "/stock", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var responseErros map[string]string

		err = json.Unmarshal(responseBody, &responseErros)

		if err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Equal(t, responseErros["quantity"], "cannot be negative value")
	})

	t.Run("Should return status code 404 and erro body product not found", func(t *testing.T) {
		invalidStockDto := &dto.StockDto{
			Product_id: 2,
			Quantity: 10,
		}

		body, _  := json.Marshal(invalidStockDto)

		req, _ := http.NewRequest("POST", url + "/stock", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var responseErros map[string]string

		err = json.Unmarshal(responseBody, &responseErros)

		if err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
		assert.Equal(t, fmt.Sprintf("product with id %d not found", invalidStockDto.Product_id), responseErros["error"])
	})

	t.Run("Should return status code 200 and stock body", func(t *testing.T) {
		productId := "1"

		req, _ := http.NewRequest("GET", url + "/stock/" + productId, nil)

		client := &http.Client{}
		res, err := client.Do(req)
		
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var responseStock domain.Stock

		err = json.Unmarshal(responseBody, &responseStock)

		if err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, stockDto.Product_id, responseStock.Product_id)
		assert.Equal(t, stockDto.Quantity, responseStock.Quantity)
	})

	t.Run("Should return status code 404 and error body stock", func(t *testing.T) {
		productId := "2"

		req, _ := http.NewRequest("GET", url + "/stock/" + productId, nil)

		client := &http.Client{}
		res, err := client.Do(req)
		
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusNotFound, res.StatusCode)
		assert.Equal(t, fmt.Sprintf("stock with product_id %s not found\n", productId), string(responseBody))
	})
}