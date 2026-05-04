package driven

import (
	"context"

	"github.com/tip-platform/tip-ranktrack/internal/player/domain"
)

type PlayerRepository interface {
	Save(ctx context.Context, player domain.Player) error
	FindByID(ctx context.Context, id uint32) (domain.Player, error)
}
