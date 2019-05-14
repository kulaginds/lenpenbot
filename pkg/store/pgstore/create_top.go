package pgstore

import (
	"time"

	"github.com/kulaginds/lenpenbot/pkg/store"
)

func (s *PGStore) CreateTop(chatID int64, top string) error {
	_, err := s.db.Exec(
		`INSERT INTO top(chat_id, type, message, updated, created)
							VALUES ($1, $2, $3, $4, $4)`,
		chatID,
		store.TopTypeAll,
		top,
		time.Now().UTC(),
	)

	return err
}
