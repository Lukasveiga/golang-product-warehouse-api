package postgre

import (
	"database/sql"
	"log"
	"product-warehouse/internal/domain"
)

type PostgreProductRepository struct {
	db *sql.DB
}

func NewPostgreProductRepository(db *sql.DB) *PostgreProductRepository {
	return &PostgreProductRepository{
		db: db,
	}
}

func (pr *PostgreProductRepository) AddProduct(product *domain.Product) *domain.Product {
	var savedProduct domain.Product

	query := "INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING *"

	row := pr.db.QueryRow(query, product.Name, product.Description, product.Price)
	err := row.Scan(&savedProduct.Id, &savedProduct.Name, &savedProduct.Description, &savedProduct.Price)

	if err != nil {
		log.Fatal(err)
	}

	return &savedProduct
}

func (pr *PostgreProductRepository) FindProductById(id int) *domain.Product {
	var product domain.Product

	query := "SELECT * FROM products WHERE id = $1"

	row := pr.db.QueryRow(query, id)
	err := row.Scan(&product.Id, &product.Name, &product.Description, &product.Price)

	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		log.Fatal(err)
	}

	return &product
}