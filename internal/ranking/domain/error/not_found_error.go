// Package error defines typed domain errors for the ranking context.
// These errors are returned by the domain layer and mapped to
// appropriate gRPC status codes by the infrastructure drivers.
//
// Using typed errors instead of plain strings allows callers to
// distinguish between validation failures, not found scenarios,
// and unexpected internal errors without parsing error messages.
package error

import "fmt"

// NotFoundError represents a resource not found error.
// Used when a requested entity does not exist.
type NotFoundError struct {
	Entity string
	ID     string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID '%s' not found", e.Entity, e.ID)
}
