package controller

import (
	"encoding/json"
	"net/http"
	"product-warehouse/internal/usecase/dto"
	usecase "product-warehouse/internal/usecase/stock"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type StockController struct {
	createStock *usecase.CreateStockUsecase
	findStockByProductId *usecase.FindStockByProductIdUsecase
}

func NewStockController(createStock *usecase.CreateStockUsecase, findStockByProductId *usecase.FindStockByProductIdUsecase) *StockController {
	return &StockController{
		createStock: createStock,
		findStockByProductId: findStockByProductId,
	}
}

func (sc StockController) Create(res http.ResponseWriter, req *http.Request) {
	var stockDto dto.StockDto
	err := json.NewDecoder(req.Body).Decode(&stockDto)

	if err != nil {
		http.Error(res, "Decoding error", http.StatusBadRequest)
		return
	}

	newStock, err := sc.createStock.Execute(stockDto)

	if err != nil {
		http.Error(res, "Product not found", http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(newStock)
}

func (sc StockController) FindStockByProductId(res http.ResponseWriter, req *http.Request) {
	productId, err := strconv.Atoi(chi.URLParam(req, "productId"))

	if err != nil {
		http.Error(res, "Invalid productId", http.StatusBadRequest)
		return
	}

	stock, err := sc.findStockByProductId.Execute(productId)

	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(stock)
}