package domain

import (
	rankingerr "github.com/tip-platform/tip-ranktrack/internal/ranking/domain/error"
	"github.com/tip-platform/tip-ranktrack/internal/ranking/domain/model"
)

// Re-export ranking domain models.
type RankingFilter = model.RankingFilter
type TrajectoryFilter = model.TrajectoryFilter
type RankingSnapshot = model.RankingSnapshot
type TrajectoryPoint = model.TrajectoryPoint
type PlayerTrajectory = model.PlayerTrajectory

// Re-export ranking constants.
type RankingType = model.RankingType

const (
	ATP = model.ATP
	WTA = model.WTA
)

// Re-export ranking domain errors.
type InvalidPeriodError = rankingerr.InvalidPeriodError
type InvalidBestDistanceError = rankingerr.InvalidBestDistanceError
type InvalidPlayerComparisonError = rankingerr.InvalidPlayerComparisonError
type InvalidRankingLimitError = rankingerr.InvalidRankingLimitError
type InvalidFutureDateError = rankingerr.InvalidFutureDateError
type InvalidRankingTypeError = rankingerr.InvalidRankingTypeError
type NotFoundError = rankingerr.NotFoundError
