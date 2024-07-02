package postgre

import (
	"database/sql"
	"log/slog"
	"product-warehouse/internal/domain"
)

type PostgreStockRepository struct {
	db *sql.DB
}

func NewPostgreStockRepository(db *sql.DB) *PostgreStockRepository {
	return &PostgreStockRepository{
		db: db,
	}
}

func (sr *PostgreStockRepository) AddStock(stock *domain.Stock) *domain.Stock {
	var savedStock domain.Stock

	query := "INSERT INTO stocks (product_id, quantity) VALUES ($1, $2) RETURNING *"

	row := sr.db.QueryRow(query, stock.Product_id, stock.Quantity)
	err := row.Scan(&savedStock.Id, &savedStock.Product_id, &savedStock.Quantity)

	if err != nil {
		slog.Error("postgre stock repository", "method", "AddStock", "error", err)
	}

	return &savedStock
}

func (sr *PostgreStockRepository) FindStockByProductId(productId int) (*domain.Stock) {
	var stock domain.Stock

	query := "SELECT * FROM stocks WHERE product_id = $1"

	row := sr.db.QueryRow(query, productId)
	err := row.Scan(&stock.Id, &stock.Product_id, &stock.Quantity)

	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		slog.Error("postgre stock repository", "method", "FindStockByProductId", "error", err)
	}

	return &stock
}

func (sr *PostgreStockRepository) UpdateStockQuantity(stockId int, quantity int) *domain.Stock {
	var stock domain.Stock

	query := "UPDATE stocks SET quantity = $1 WHERE id = $2 RETURNING *"

	row := sr.db.QueryRow(query, quantity, stockId)
	err := row.Scan(&stock.Id, &stock.Product_id, &stock.Quantity)

	if err != nil {
		slog.Error("postgre stock repository", "method", "UpdateStockQuantity", "error", err)
	}

	return &stock
}