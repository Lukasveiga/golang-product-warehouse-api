package dto

import (
	"product-warehouse/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, stockTest, *resultTest)
}