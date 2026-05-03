package model

import err "github.com/tip-platform/tip-ranktrack/internal/player/domain/error"

type Player struct {
	ID          uint32
	APIID       uint32
	Name        string
	ShortName   string
	CountryCode string
	CountryName string
	Age         uint8
	Plays       string
	TurnedPro   uint16
}

func (p Player) Validate() error {
	if p.Name == "" {
		return err.InvalidPlayerNameError{}
	}
	if p.CountryCode == "" {
		return err.InvalidPlayerCountryCodeError{}
	}
	return nil
}
