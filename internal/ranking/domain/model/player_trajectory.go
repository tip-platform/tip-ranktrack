package model

import playerdomain "github.com/tip-platform/tip-ranktrack/internal/player/domain"

type PlayerTrajectory struct {
	Player playerdomain.Player
	Points []TrajectoryPoint
}
