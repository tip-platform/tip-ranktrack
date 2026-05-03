package error

import (
	"fmt"
	"time"
)

// InvalidFutureDateError represents a date in the future.
type InvalidFutureDateError struct {
	Value time.Time
}

func (e InvalidFutureDateError) Error() string {
	return fmt.Sprintf("invalid date: '%s' must not be in the future", e.Value.Format(time.RFC3339))
}
