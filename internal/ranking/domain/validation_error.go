// Package domain defines the core business entities for the rankings context.
// ValidationError represents a domain validation error.
// Used when input data does not meet business rules.
package domain

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error: field '%s' - %s", e.Field, e.Message)
}
