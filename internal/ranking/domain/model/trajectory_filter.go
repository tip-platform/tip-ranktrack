package model

import (
	"time"

	domainerr "github.com/tip-platform/tip-ranktrack/internal/ranking/domain/error"
)

type TrajectoryFilter struct {
	PlayerID  uint32
	Period    string
	Player2ID *uint32
	Type      RankingType
	Date      time.Time
}

func (f TrajectoryFilter) Validate() error {
	if f.Type != ATP && f.Type != WTA {
		return domainerr.InvalidRankingTypeError{Value: string(f.Type)}
	}
	if !ValidPeriod[f.Period] {
		return domainerr.InvalidPeriodError{Value: f.Period}
	}
	if f.Player2ID != nil && *f.Player2ID == f.PlayerID {
		return domainerr.InvalidPlayerComparisonError{PlayerID: f.PlayerID}
	}
	if f.Date.After(time.Now()) {
		return domainerr.InvalidFutureDateError{Value: f.Date}
	}
	return nil
}
