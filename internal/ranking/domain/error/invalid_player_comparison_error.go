package error

import "fmt"

// InvalidPlayerComparisonError represents an invalid comparison player selection.
type InvalidPlayerComparisonError struct {
	PlayerID uint32
}

func (e InvalidPlayerComparisonError) Error() string {
	return fmt.Sprintf("invalid player comparison: player2_id must differ from player_id '%d'", e.PlayerID)
}
