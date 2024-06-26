package postgre

import (
	"product-warehouse/config"
	"product-warehouse/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgreProductRepository(t *testing.T) {

	dbClient := config.SetupDbConnection()

	postgreProductRepository := NewPostgreProductRepository(dbClient)

	defer dbClient.Close()

	t.Run("Should successfully connect to the database", func(t *testing.T) {
		assert.NotNil(t, &dbClient)
	})

	t.Run("Should save product and then return it", func(t *testing.T) {
		product := &domain.Product{
			Name: "Product Test Name",
			Description: "Product Test Description",
			Price: 10.0,
		}

		savedProduct := postgreProductRepository.AddProduct(product)

		assert.NotNil(t, savedProduct)

		product.Id = savedProduct.Id

		assert.Equal(t, product, savedProduct)
	})

	t.Run("Should found product by id", func(t *testing.T) {
		productId := 1

		product := postgreProductRepository.FindProductById(productId)

		assert.NotNil(t, product)
		assert.Equal(t, productId, product.Id)
	})

	t.Run("Should return nil if product was not found", func(t *testing.T) {
		productId := 2

		product := postgreProductRepository.FindProductById(productId)

		assert.Nil(t, product)
	})
}