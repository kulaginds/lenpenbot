package pgstore

import "github.com/kulaginds/lenpenbot/pkg/store"

func (s *PGStore) HasTop(chatID int64) (bool, error) {
	rows, err := s.db.Query(
		`SELECT NULL FROM "top" WHERE chat_id = $1 AND type = $2`,
		chatID,
		store.TopTypeAll,
	)
	if err != nil {
		return false, err
	}

	return rows.Next(), nil
}
