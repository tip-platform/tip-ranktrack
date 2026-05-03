// Package model defines the core business entities for the rankings
// repository interface that driven adapters must implement.
package model

import (
	"time"

	playerdomain "github.com/tip-platform/tip-ranktrack/internal/player/domain"
	modelerr "github.com/tip-platform/tip-ranktrack/internal/ranking/domain/error"
)

type RankingType string

const (
	ATP RankingType = "ATP"
	WTA RankingType = "WTA"
)

type RankingSnapshot struct {
	ID                                        uint32
	Player                                    playerdomain.Player
	RankingType                               RankingType
	SnapshotDate                              time.Time
	Momentum, Volatility                      float64
	BestPosition, DistanceFromBest            uint16
	Points, HighestPoints, LowestPoints       uint16
	Position, HighestPosition, LowestPosition uint16
}

func (s RankingSnapshot) Validate() error {
	if s.DistanceFromBest > s.BestPosition {
		return modelerr.InvalidBestDistanceError{
			DistanceFromBest: s.DistanceFromBest,
			BestPosition:     s.BestPosition,
		}
	}
	return nil
}
