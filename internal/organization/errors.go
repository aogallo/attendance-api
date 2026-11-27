package organization

import "errors"

var (
	ErrForbidden   = errors.New("platform authority required")
	ErrUnavailable = errors.New("persistence unavailable")
)

type FieldError struct {
	Field string `json:"field"`
	Code  string `json:"code"`
}
type ValidationError struct{ Fields []FieldError }

func (e *ValidationError) Error() string { return "validation failed" }
