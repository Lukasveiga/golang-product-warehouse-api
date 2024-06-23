package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"product-warehouse/cmd/api/utils"
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
		http.Error(res, "Decoding error", http.StatusBadRequest)
		return
	}

	newStock, errs := sc.createStock.Execute(stockDto)

	if errs != nil {
		errsString, err := utils.ErrorFormatter(errs)

		if err != nil {
			log.Print(err)
			http.Error(res, "Internal error", http.StatusInternalServerError)
			return
		}

		_, key := errs["error"]

		if key {
			http.Error(res, errsString, http.StatusNotFound)
			return
		}

		http.Error(res, errsString, http.StatusBadRequest)
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

	stock, errs := sc.findStockByProductId.Execute(productId)

	if errs != nil {
		http.Error(res, errs["error"].Error(), http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(stock)
}