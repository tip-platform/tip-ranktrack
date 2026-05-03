package adapter

import (
	"github.com/tip-platform/tip-ranktrack/internal/player/domain/model"
	"github.com/tip-platform/tip-ranktrack/internal/player/infra/persistence/entity"
)

func ToDomain(r entity.PlayerRecord) model.Player {
	return model.Player{
		ID:          r.ID,
		APIID:       r.APIID,
		Name:        r.Name,
		ShortName:   r.ShortName,
		CountryCode: r.CountryCode,
		CountryName: r.CountryName,
		Age:         r.Age,
		Plays:       r.Plays,
		TurnedPro:   r.TurnedPro,
	}
}

func ToRecord(p model.Player) entity.PlayerRecord {
	return entity.PlayerRecord{
		ID:           p.ID,
		APIID:       p.APIID,
		Name:         p.Name,
		SHORT_Name:   p.ShortName,
		CountryCode: p.CountryCode,
		COUNTRY_Name: p.CountryName,
		Age:          p.Age,
		Plays:        p.Plays,
		TurnedPro:   p.TurnedPro,
	}
}
