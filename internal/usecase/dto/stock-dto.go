package dto

import (
	"product-warehouse/internal/domain"
)

type StockDto struct {
	Product_id int `json:"product_id"`
	Quantity int `json:"quantity"`
}

func NewStock(stockDto *StockDto) (*domain.Stock, error) {
	s := &domain.Stock{
		Product_id: stockDto.Product_id,
		Quantity: stockDto.Quantity,
	}

	errs := s.Validate()

	if errs != nil {
		return nil, errs
	}

	return s, nil
}
