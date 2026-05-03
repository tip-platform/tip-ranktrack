package port

import (
	"github.com/tip-platform/tip-ranktrack/internal/player/application/port/driven"
	"github.com/tip-platform/tip-ranktrack/internal/player/application/port/driver"
)

// PlayerService re-exports the driver port for the application layer.
type PlayerService = driver.PlayerService

// PlayerRepository re-exports the driven port for the application layer.
type PlayerRepository = driven.PlayerRepository
