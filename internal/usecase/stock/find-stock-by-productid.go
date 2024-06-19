package usecase

import (
	"fmt"
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
)

type FindStockByProductIdUsecase struct {
	repo port.StockRepository
}

func NewFindStockByProductIdUsecase(repo port.StockRepository) * FindStockByProductIdUsecase {
	return &FindStockByProductIdUsecase{
		repo: repo,
	}
}

func (sc FindStockByProductIdUsecase) Execute(productId int) (*domain.Stock, error) {
	stock := sc.repo.FindStockByProductId(productId)

	if stock == nil {
		return nil, fmt.Errorf("product with id %d not found", productId)
	}
	return stock, nil
}