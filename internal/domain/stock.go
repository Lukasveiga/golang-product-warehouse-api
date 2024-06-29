package domain

import (
	"product-warehouse/internal/shared"
)

type Stock struct {
	Id int `json:"id"`
	Product_id int `json:"product_id"`
	Quantity int `json:"quantity"`
}

func (s *Stock) Validate() error {
	validationError := &shared.ValidationError{
		Errors: make(map[string]string),
	}

	if s.Quantity < 0 {
		validationError.AddError("quantity", "cannot be negative value")
	}

	if validationError.HasErrors() {
		return validationError
	}

	return nil
}