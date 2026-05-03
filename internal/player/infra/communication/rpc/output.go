package rpc

import (
	"github.com/tip-platform/tip-ranktrack/internal/player/domain"
	pb "github.com/tip-platform/tip-ranktrack/proto"
)

func ToPlayerInfoProto(p domain.Player) *pb.PlayerInfo {
	return &pb.PlayerInfo{
		Id:          int32(p.ID),
		Name:        p.Name,
		CountryCode: p.CountryCode,
		Age:         int32(p.Age),
	}
}
