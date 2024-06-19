package dto

import (
	"product-warehouse/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStockDto(t *testing.T) {
	t.Run("NewStock Success", func(t *testing.T) {
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
	})
}