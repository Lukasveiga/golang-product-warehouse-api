package dto

import (
	"product-warehouse/internal/domain"
	"product-warehouse/internal/shared"
)

type StockDto struct {
	Product_id int `json:"product_id"`
	Quantity int `json:"quantity"`
}

func NewStock(stockDto *StockDto) (*domain.Stock, shared.ErrorMap) {
	s := &domain.Stock{
		Product_id: stockDto.Product_id,
		Quantity: stockDto.Quantity,
	}

	errs := s.Validate()

	if len(errs) != 0 {
		return nil, errs
	}

	return s, nil
}
