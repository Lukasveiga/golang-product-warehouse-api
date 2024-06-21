package dto

import (
	"errors"
	"product-warehouse/internal/domain"
	"product-warehouse/internal/shared"
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
	
		resultTest, err := NewStock(&stockStoTest)
	
		assert.Nil(t, err)
		assert.Equal(t, stockTest, *resultTest)
	})

	t.Run("NewStock Invalid Input", func(t *testing.T) {
		stockStoTest := StockDto{
			Product_id: 1,
			Quantity: -1,
		}

		expectedError := shared.ErrorMap{
			"quantity": errors.New("quantity cannot be negative value"),
		}

		resultTest, err := NewStock(&stockStoTest)

		assert.Nil(t, resultTest)
		assert.Equal(t, expectedError, err)
	})
}