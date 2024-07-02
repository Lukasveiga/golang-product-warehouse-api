package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"product-warehouse/internal/shared"
	"product-warehouse/internal/usecase/dto"
	usecase "product-warehouse/internal/usecase/product"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductController struct {
	createProductUsecase *usecase.CreateProductUsecase
	findProductByIdUsecase *usecase.FindProductByIdUsecase
}

func NewProductController(createProductUseCase *usecase.CreateProductUsecase, findProductByIdUsecase *usecase.FindProductByIdUsecase) *ProductController {
	return &ProductController{
		createProductUsecase: createProductUseCase,
		findProductByIdUsecase: findProductByIdUsecase,
	}
}

func (pc ProductController) Create(res http.ResponseWriter, req *http.Request) {
	var productDto *dto.ProductDto
	err := json.NewDecoder(req.Body).Decode(&productDto)

	if err != nil {
		http.Error(res, "Decoding Error", http.StatusBadRequest)
		return
	}

	newProduct, err := pc.createProductUsecase.Execute(productDto)

	if err != nil {
		if ve, ok := err.(*shared.ValidationError); ok {
			jsonData, err := json.Marshal(ve.Errors)

			if err != nil {
				slog.Error("product controller", "method", "Create", "error", err)
				http.Error(res, "Internal error", http.StatusInternalServerError)
				return
			}

			http.Error(res, string(jsonData), http.StatusBadRequest)
			return

		} else {
			slog.Error("product controller", "method", "Create", "error", err)
			http.Error(res, "Internal error", http.StatusInternalServerError)
			return
		}
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(newProduct)
}

func (pc ProductController) FindById(res http.ResponseWriter, req *http.Request) {
	productId, err := strconv.Atoi(chi.URLParam(req, "id"))
	
	if err != nil {
		http.Error(res, "Invalid id", http.StatusBadRequest)
		return
	}

	product, err := pc.findProductByIdUsecase.Execute(productId)

	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(product)
}