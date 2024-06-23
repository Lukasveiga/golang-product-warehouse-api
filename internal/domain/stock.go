package domain

import (
	"errors"
	"product-warehouse/internal/shared"
)

type Stock struct {
	Id int `json:"id"`
	Product_id int `json:"product_id"`
	Quantity int `json:"quantity"`
}

func (s *Stock) Validate() shared.ErrorMap {
	errs := make(shared.ErrorMap)

	if s.Quantity < 0 {
		errs["quantity"] = errors.New("cannot be negative value")
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}