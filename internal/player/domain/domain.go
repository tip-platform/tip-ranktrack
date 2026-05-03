package domain

import (
	err "github.com/tip-platform/tip-ranktrack/internal/player/domain/error"
	"github.com/tip-platform/tip-ranktrack/internal/player/domain/model"
)

// Re-export player domain models.
type Player = model.Player

// Re-export player domain errors.
type InvalidPlayerNameError = err.InvalidPlayerNameError
type InvalidPlayerCountryCodeError = err.InvalidPlayerCountryCodeError
