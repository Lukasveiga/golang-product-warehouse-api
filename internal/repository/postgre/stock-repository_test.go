package postgre

import (
	"product-warehouse/config"
	"product-warehouse/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgreStockRepository(t *testing.T) {
	dbClient := config.SetupDbConnection()

	postgreStockRepository := NewPostgreStockRepository(dbClient)
	postgreProductRepository := NewPostgreProductRepository(dbClient)

	product := postgreProductRepository.AddProduct(&domain.Product{
		Name: "Product Test Name",
		Description: "Product Test Description",
		Price: 10.0,
	})

	defer dbClient.Close()

	t.Run("Should successfully connect to the database", func(t *testing.T) {
		assert.NotNil(t, &dbClient)
	})

	t.Run("Should save product and then return it", func(t *testing.T) {
		stock := &domain.Stock{
			Product_id: product.Id,
			Quantity: 10,
		}

		savedStock := postgreStockRepository.AddStock(stock)

		assert.NotNil(t, savedStock)

		stock.Id = savedStock.Id

		assert.Equal(t, stock, savedStock)
	})

	t.Run("Should found stock by product id", func(t *testing.T) {
		stock := postgreStockRepository.FindStockByProductId(product.Id)

		assert.NotNil(t, stock)
		assert.Equal(t, product.Id, stock.Product_id)
	})

	t.Run("Should return nil if stock was not found", func(t *testing.T) {
		productId := 2

		stock := postgreStockRepository.FindStockByProductId(productId)

		assert.Nil(t, stock)
	})

	t.Run("Should update stock quantity", func(t *testing.T) {
		stock := postgreStockRepository.FindStockByProductId(product.Id)

		quantity := 20

		updatedStock := postgreStockRepository.UpdateStockQuantity(stock.Id, quantity)

		assert.NotNil(t, updatedStock)
		assert.Equal(t, quantity, updatedStock.Quantity)
	})
}