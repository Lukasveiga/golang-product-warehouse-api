package port

import (
	"product-warehouse/internal/domain"
)

type ProductRepository interface {
	AddProduct(product *domain.Product) *domain.Product
	FindProductById(id int) (*domain.Product, error)
}