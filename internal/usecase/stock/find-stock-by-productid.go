package usecase

import (
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

func (sc FindStockByProductIdUsecase) Execute(productId int) (*domain.Stock, error) {
	stock := sc.repo.FindStockByProductId(productId)

	if stock == nil {
		return nil, &shared.NotFoundError{
			Object: "product",
			Id: productId,
		} 
	}
	return stock, nil
}