package rpc

import (
	"context"

	"github.com/tip-platform/tip-ranktrack/internal/player/application/port/driver"
	pb "github.com/tip-platform/tip-ranktrack/proto"
)

type PlayerHandler struct {
	pb.UnimplementedRankTrackServiceServer
	playerService driver.PlayerService
}

func NewPlayerHandler(playerService driver.PlayerService) *PlayerHandler {
	return &PlayerHandler{
		playerService: playerService,
	}
}

func (h *PlayerHandler) GetPlayer(ctx context.Context, req *pb.GetPlayerRequest) (*pb.GetPlayerResponse, error) {
	player, err := h.playerService.GetPlayer(ctx, uint32(req.PlayerId))
	if err != nil {
		return nil, err
	}

	return &pb.GetPlayerResponse{
		Player: ToPlayerInfoProto(player),
		// Other fields (momentum, volatility, etc.) are typically handled by other domain services or combined at application level.
		// For now, we return the player info as per infrastructure scope.
	}, nil
}
