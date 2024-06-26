package inMemory

import (
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

func (db *InMemoryStockRepository) FindStockByProductId(productId int) (*domain.Stock) {
	result := []*domain.Stock{}

	for i := range db.inMemoryStockDb {
		if db.inMemoryStockDb[i].Product_id == productId {
			result = append(result, &db.inMemoryStockDb[i])
			break
		}
	}

	if len(result) == 0 {
		return nil
	}

	return result[0]
}

func (db *InMemoryStockRepository) UpdateStockQuantity(stockId int, quantity int) *domain.Stock {
	var stock *domain.Stock
	for i := range db.inMemoryStockDb {
		if db.inMemoryStockDb[i].Id == stockId {
			stock = &db.inMemoryStockDb[i]
			stock.Quantity = quantity
			break
		}
	}
	return stock
}