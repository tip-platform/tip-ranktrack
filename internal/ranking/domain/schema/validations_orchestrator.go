// Package schema provides domain validation logic for input filters.
// It defines the Validatable interface and orchestrates validation
// of any domain filter before it reaches the application layer.
//
// Any type that implements Validate() error satisfies Validatable
// and can be validated through this package without schema knowing
// the concrete type.
package schema

import "github.com/tip-platform/tip-ranktrack/internal/ranking/domain/model"

func Validate(v model.Validator) error {
	return v.Validate()
}
