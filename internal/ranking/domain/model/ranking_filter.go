package model

import (
	"time"

	domainerr "github.com/tip-platform/tip-ranktrack/internal/ranking/domain/error"
)

type RankingFilter struct {
	Type    RankingType
	Limit   uint16
	Country *string
	Date    time.Time
	Period  string
}

var ValidPeriod = map[string]bool{
	"1w": true, "2w": true, "1m": true,
	"3m": true, "6m": true, "1y": true,
	"2y": true, "5y": true,
}

func (m RankingFilter) Validate() error {
	if m.Type != ATP && m.Type != WTA {
		return domainerr.InvalidRankingTypeError{Value: string(m.Type)}
	}
	if m.Type == ATP && m.Limit > 2000 {
		return domainerr.InvalidRankingLimitError{Limit: m.Limit, Max: 2000, Type: string(m.Type)}
	}
	if !ValidPeriod[m.Period] {
		return domainerr.InvalidPeriodError{Value: m.Period}
	}
	if m.Date.After(time.Now()) {
		return domainerr.InvalidFutureDateError{Value: m.Date}
	}
	return nil
}
