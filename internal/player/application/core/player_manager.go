package core

import (
	"context"

	"github.com/tip-platform/tip-ranktrack/internal/player/application/port"
	"github.com/tip-platform/tip-ranktrack/internal/player/domain"
)

type playerManager struct {
	repo port.PlayerRepository
}

func (m *playerManager) CreatePlayer(ctx context.Context, player domain.Player) error {
	if err := player.Validate(); err != nil {
		return err
	}

	return m.repo.Save(ctx, player)
}

func (m *playerManager) GetPlayer(ctx context.Context, id uint32) (domain.Player, error) {
	return m.repo.FindByID(ctx, id)
}
