package usecase

import (
	"fmt"
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
	"product-warehouse/internal/shared"
)

type FindStockByProductIdUsecase struct {
	repo port.StockRepository
}

func NewFindStockByProductIdUsecase(repo port.StockRepository) * FindStockByProductIdUsecase {
	return &FindStockByProductIdUsecase{
		repo: repo,
	}
}

func (sc FindStockByProductIdUsecase) Execute(productId int) (*domain.Stock, shared.ErrorMap) {
	stock := sc.repo.FindStockByProductId(productId)

	if stock == nil {
		return nil, shared.ErrorMap{
			"error": fmt.Errorf("stock with product_id %d not found", productId),
		}
	}
	return stock, nil
}