package port

import "product-warehouse/internal/domain"

type StockRepository interface {
	AddStock(stock *domain.Stock) (*domain.Stock)
	FindStockByProductId(productId int) (*domain.Stock, error)
	UpdateStockQuantity(stockId int, quantity int) *domain.Stock
}