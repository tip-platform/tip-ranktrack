package persistence

import (
	"database/sql"
	"fmt"

	"github.com/tip-platform/tip-ranktrack/internal/player/domain/model"
	"github.com/tip-platform/tip-ranktrack/internal/player/infra/persistence/adapter"
	"github.com/tip-platform/tip-ranktrack/internal/player/infra/persistence/entity"
	persistence_sql "github.com/tip-platform/tip-ranktrack/internal/player/infra/persistence/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Save(player model.Player) error {
	record := adapter.ToRecord(player)

	_, err := s.db.Exec(persistence_sql.INSERT_PLAYER,
		record.APIID,
		record.Name,
		record.ShortName,
		record.CountryCode,
		record.CountryName,
		record.Age,
		record.Plays,
		record.TurnedPro,
	)

	if err != nil {
		return fmt.Errorf("error saving player: %w", err)
	}

	return nil
}

func (s *Store) GetByID(id uint32) (model.Player, error) {
	var record entity.PlayerRecord

	err := s.db.QueryRow(persistence_sql.SELECT_PLAYER_BY_ID, id).Scan(
		&record.ID,
		&record.APIID,
		&record.Name,
		&record.ShortName,
		&record.CountryCode,
		&record.CountryName,
		&record.Age,
		&record.Plays,
		&record.TurnedPro,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return model.Player{}, fmt.Errorf("player not found: %w", err)
		}
		return model.Player{}, fmt.Errorf("error getting player: %w", err)
	}

	return adapter.ToDomain(record), nil
}
