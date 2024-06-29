package domain

import (
	"product-warehouse/internal/shared"
)

type Product struct {
	Id int `json:"id"`
	Name string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Price float64 `json:"price" db:"price"`
}

func (p *Product) Validate() error {
	validationError := &shared.ValidationError{
		Errors: make(map[string]string),
	}

	if p.Name == "" {
		validationError.AddError("name", "cannot be empty")
	}

	if p.Description == "" {
		validationError.AddError("description", "cannot be empty")
	}

	if p.Price <= 0 {
		validationError.AddError("price", "must be greater than zero")
	}

	if validationError.HasErrors() {
		return validationError
	}

	return nil
}

