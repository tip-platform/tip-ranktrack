package driver

import (
	"context"

	"github.com/tip-platform/tip-ranktrack/internal/player/domain"
)

type PlayerService interface {
	CreatePlayer(ctx context.Context, player domain.Player) error
	GetPlayer(ctx context.Context, id uint32) (domain.Player, error)
}
