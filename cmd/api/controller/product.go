package controller

import (
	"encoding/json"
	"net/http"
	"product-warehouse/internal/domain/dto"
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
	var productDto dto.ProductDto
	err := json.NewDecoder(req.Body).Decode(&productDto)

	if err != nil {
		http.Error(res, "Decoding error", http.StatusBadRequest)
		return
	}

	newProduct := pc.createProductUsecase.Execute(productDto)

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