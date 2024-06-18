package inMemory

import (
	"product-warehouse/internal/domain"
)

type InMemoryProductRepository struct {
	inMemoryProductDb []domain.Product
}

func NewInMemoryProductRepository() * InMemoryProductRepository {
	return &InMemoryProductRepository{
		inMemoryProductDb: []domain.Product{},
	}
}

func (db *InMemoryProductRepository) AddProduct(product *domain.Product) *domain.Product {
	product.Id = len(db.inMemoryProductDb) + 1
	db.inMemoryProductDb = append(db.inMemoryProductDb, *product)
	return product;
}

func (db *InMemoryProductRepository) FindProductById(id int) *domain.Product {
	result := []*domain.Product{}

	for i := range db.inMemoryProductDb {
		if db.inMemoryProductDb[i].Id == id {
			result = append(result, &db.inMemoryProductDb[i])
			break
		}
	}

	if len(result) == 0 {
		return nil
	}

	return result[0]
}