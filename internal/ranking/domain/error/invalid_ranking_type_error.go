package error

import "fmt"

// InvalidRankingTypeError represents an invalid ranking type error.
type InvalidRankingTypeError struct {
	Value string
}

func (e InvalidRankingTypeError) Error() string {
	return fmt.Sprintf("invalid ranking type: '%s' must be ATP or WTA", e.Value)
}
