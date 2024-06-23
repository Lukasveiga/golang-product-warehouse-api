package controller

import (
	"product-warehouse/internal/domain"
	usecase "product-warehouse/internal/usecase/stock"
)

type MockStock struct {
	Stock *domain.Stock
}

func (m *MockStock) AddStock(stock *domain.Stock) *domain.Stock {
	return m.Stock
}

func (m *MockStock) FindStockByProductId(productId int) *domain.Stock {
	return m.Stock
}

func (m *MockStock) UpdateStockQuantity(stockId int, quantity int) *domain.Stock {
	return m.Stock
}

func StockControllerSetup(stockMock *MockStock, productMock *MockProduct) *StockController {
	createStockUsecase := usecase.NewCreateStockUsecase(stockMock, productMock)
	findStockByProductIdUsecase := usecase.NewFindStockByProductIdUsecase(stockMock)
	return NewStockController(createStockUsecase, findStockByProductIdUsecase)
}