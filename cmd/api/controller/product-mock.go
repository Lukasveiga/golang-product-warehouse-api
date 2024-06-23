package controller

import (
	"product-warehouse/internal/domain"
	usecase "product-warehouse/internal/usecase/product"
)

type MockProduct struct {
	Product *domain.Product
}

func(m *MockProduct) AddProduct(product *domain.Product) *domain.Product {
	return m.Product
}

func (m *MockProduct) FindProductById(id int) (*domain.Product) {
	return m.Product
}

func ProductControllerSetup(mock *MockProduct) *ProductController {
	createProductUsecase := usecase.NewCreateProductUsecase(mock)
	findProductUsecase := usecase.NewFindProductByIdUsecase(mock)
	return NewProductController(createProductUsecase, findProductUsecase)
}