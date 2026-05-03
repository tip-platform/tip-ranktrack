// Package schema provides domain validation logic for input filters.
// It defines the Validatable interface and orchestrates validation
// of any domain filter before it reaches the application layer.
//
// Any type that implements Validate() error satisfies Validatable
// and can be validated through this package without schema knowing
// the concrete type.
package schema

// Validatable is implemented by any domain type that can be validated.
type Validatable interface {
	Validate() error
}

func Validate(v Validatable) error {
	return v.Validate()
}
