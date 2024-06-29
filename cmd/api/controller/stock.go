package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"product-warehouse/internal/shared"
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
	var stockDto *dto.StockDto
	err := json.NewDecoder(req.Body).Decode(&stockDto)

	if err != nil {
		http.Error(res, "Decoding Error", http.StatusBadRequest)
		return
	}

	newStock, err := sc.createStock.Execute(stockDto)

	if err != nil {
		if nf, ok := err.(*shared.NotFoundError); ok {
			http.Error(res, nf.Error(), http.StatusNotFound)
			return
			
		} else if ve, ok := err.(*shared.ValidationError); ok {
			jsonData, err := json.Marshal(ve.Errors)

			if err != nil {
				log.Print(err)
				http.Error(res, "Internal error", http.StatusInternalServerError)
				return
			}

			http.Error(res, string(jsonData), http.StatusBadRequest)
			return
			
		} else {
			log.Print(err)
			http.Error(res, "Internal error", http.StatusInternalServerError)
			return
		}
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