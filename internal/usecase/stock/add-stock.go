package usecase

import (
	"fmt"
	"product-warehouse/internal/domain"
	port "product-warehouse/internal/port/repository"
	"product-warehouse/internal/shared"
	"product-warehouse/internal/usecase/dto"
)

type CreateStockUsecase struct {
	stockRepo port.StockRepository
	productRepo port.ProductRepository
}

func NewCreateStockUsecase(stockRepo port.StockRepository, productRepo port.ProductRepository) *CreateStockUsecase {
	return &CreateStockUsecase{
		stockRepo: stockRepo,
		productRepo: productRepo,
	}
}

func (sc CreateStockUsecase) Execute(stockDto *dto.StockDto) (*domain.Stock, shared.ErrorMap) {
	product := sc.productRepo.FindProductById(stockDto.Product_id)

	if product == nil {
		return nil, shared.ErrorMap{
			"error": fmt.Errorf("product with id %d not found", stockDto.Product_id),
		}
	}

	stock := sc.stockRepo.FindStockByProductId(stockDto.Product_id)

	s, errs := dto.NewStock(stockDto)

	if errs != nil {
		return nil, errs
	}

	if stock != nil {
		return sc.stockRepo.UpdateStockQuantity(stock.Id, stockDto.Quantity), nil
	}

	return sc.stockRepo.AddStock(s), nil
}