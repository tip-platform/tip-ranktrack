package error

import "fmt"

// InvalidPeriodError represents an invalid period error.
type InvalidPeriodError struct {
	Value string
}

func (e InvalidPeriodError) Error() string {
	return fmt.Sprintf("invalid period: '%s' must be one of 1w, 2w, 1m, 3m, 6m, 1y, 2y, 5y", e.Value)
}
