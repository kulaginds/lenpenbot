package pgstore

import (
	"time"

	"github.com/kulaginds/lenpenbot/pkg/store"
)

func (s *PGStore) UpdateTop(chatID int64, top string) error {
	_, err := s.db.Exec(
		`UPDATE
					"top"
				SET
					message = $3,
				    updated = $4
				WHERE
					chat_id = $1
					AND type = $2`,
		chatID,
		store.TopTypeAll,
		top,
		time.Now().UTC(),
	)

	return err
}
