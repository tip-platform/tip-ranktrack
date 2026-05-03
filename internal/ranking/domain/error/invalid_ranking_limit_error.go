package error

import "fmt"

// InvalidRankingLimitError represents an invalid ranking limit.
type InvalidRankingLimitError struct {
	Limit uint16
	Max   uint16
	Type  string
}

func (e InvalidRankingLimitError) Error() string {
	return fmt.Sprintf("invalid ranking limit: '%d' must be <= %d for %s", e.Limit, e.Max, e.Type)
}
