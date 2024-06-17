package inMemory

import (
	"fmt"
	"product-warehouse/internal/domain"
)

type InMemoryStockRepository struct {
	inMemoryStockDb []domain.Stock
}

func NewInMemoryStockRepository() *InMemoryStockRepository {
	return &InMemoryStockRepository{
		inMemoryStockDb: []domain.Stock{},
	}
}

func (db *InMemoryStockRepository) AddStock(stock *domain.Stock) *domain.Stock {
	stock.Id = len(db.inMemoryStockDb) + 1
	db.inMemoryStockDb = append(db.inMemoryStockDb, *stock)
	return stock;
}

func (db *InMemoryStockRepository) FindStockByProductId(productId int) (*domain.Stock, error) {
	result := []*domain.Stock{}

	for i := range db.inMemoryStockDb {
		if db.inMemoryStockDb[i].Product_id == productId {
			result = append(result, &db.inMemoryStockDb[i])
			break
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("stock with productId %d not found", productId)
	}

	return result[0], nil
}

func (db *InMemoryStockRepository) UpdateStockQuantity(stockId int, quantity int) *domain.Stock {
	var stock *domain.Stock
	for i := range db.inMemoryStockDb {
		if db.inMemoryStockDb[i].Id == stockId {
			stock = &db.inMemoryStockDb[i]
			stock.Quantity += quantity
			break
		}
	}
	return stock
}