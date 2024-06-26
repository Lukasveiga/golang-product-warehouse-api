package domain

import (
	"errors"
	"product-warehouse/internal/shared"
)

type Product struct {
	Id int `json:"id"`
	Name string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Price float64 `json:"price" db:"price"`
}

func (p *Product) Validate() shared.ErrorMap {
	errs := make(shared.ErrorMap)

	if p.Name == "" {
		errs["name"] = errors.New("cannot be empty")
	}

	if p.Description == "" {
		errs["description"] = errors.New("cannot be empty")
	}

	if p.Price <= 0 {
		errs["price"] = errors.New("must be greater than zero")
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

