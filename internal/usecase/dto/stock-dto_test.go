package dto

import (
	"product-warehouse/internal/domain"
	"reflect"
	"testing"
)

func TestStockDto_Success_NewStock(t *testing.T) {
	stockTest := domain.Stock{
		Product_id: 1,
		Quantity: 20,
	}

	stockStoTest := StockDto{
		stockTest.Product_id,
		stockTest.Quantity,
	}

	resultTest := NewStock(&stockStoTest)

	if !reflect.DeepEqual(*resultTest, stockTest) {
		t.Errorf("Product body expected: %v, got: %v", stockTest, resultTest)
	}
}