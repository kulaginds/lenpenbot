package pgstore

import (
	"time"

	"github.com/kulaginds/lenpenbot/pkg/store"
)

func (s *PGStore) CreateToday(chatID int64, today time.Time, todayTop string) error {
	_, err := s.db.Exec(
		`INSERT INTO top(chat_id, type, message, updated, created)
							VALUES ($1, $2, $3, $4, $5)`,
		chatID,
		store.TopTypeToday,
		todayTop,
		today,
		today.Truncate(24 * time.Hour),
	)

	return err
}
