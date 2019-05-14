package pgstore

import (
	"time"

	"github.com/kulaginds/lenpenbot/pkg/store"
)

func (s *PGStore) HasToday(chatID int64, updatedMin, updatedMax time.Time) (bool, error) {
	rows, err := s.db.Query(
		`SELECT NULL FROM "top" WHERE chat_id = $1 AND type = $2 AND updated BETWEEN $3 AND $4`,
		chatID,
		store.TopTypeToday,
		updatedMin,
		updatedMax,
	)
	if err != nil {
		return false, err
	}

	return rows.Next(), nil
}
