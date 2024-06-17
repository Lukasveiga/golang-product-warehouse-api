package usecase

import (
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
	stock, err := sc.repo.FindStockByProductId(productId)
	return stock, err
}