package error

import "fmt"

// InvalidBestDistanceError represents an invalid distance from best position.
type InvalidBestDistanceError struct {
	DistanceFromBest uint16
	BestPosition     uint16
}

func (e InvalidBestDistanceError) Error() string {
	return fmt.Sprintf("invalid distance_from_best: '%d' must be <= best_position '%d'", e.DistanceFromBest, e.BestPosition)
}
