package dto

import "product-warehouse/internal/domain"

type StockDto struct {
	Product_id int `json:"product_id"`
	Quantity int `json:"quantity"`
}

func NewStock(stockDto StockDto) *domain.Stock {
	return &domain.Stock{
		Product_id: stockDto.Product_id,
		Quantity: stockDto.Quantity,
	}
}
