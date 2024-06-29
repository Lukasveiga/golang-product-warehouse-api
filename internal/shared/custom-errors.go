package shared

import "fmt"

type NotFoundError struct {
	Object string
	Id int
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s not found with id %d", e.Object, e.Id)
}

type ValidationError struct {
	Errors map[string]string
}

func (e ValidationError) Error() string {
	return "validation error"
}

func (e *ValidationError) AddError(field string, err string) {
	e.Errors[field] = err
}

func (e *ValidationError) HasErrors() bool {
	return len(e.Errors) != 0;
}

