package usecase

import (
	"product-warehouse/internal/domain"
	"product-warehouse/internal/domain/dto"
	port "product-warehouse/internal/port/repository"
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

func (sc CreateStockUsecase) Execute(stockDto dto.StockDto) (*domain.Stock, error) {
	_, err := sc.productRepo.FindProductById(stockDto.Product_id)

	if err != nil {
		return nil, err
	}

	stock, _ := sc.stockRepo.FindStockByProductId(stockDto.Product_id)

	if stock != nil {
		return sc.stockRepo.UpdateStockQuantity(stock.Id, stockDto.Quantity), nil
	}

	return sc.stockRepo.AddStock(dto.NewStock(stockDto)), nil
}