package pgstore

import (
	"time"
)

func (s *PGStore) IsEnlarge(userID int, chatID int64, date time.Time) (bool, error) {
	rows, err := s.db.Query(
		`SELECT NULL FROM enlarge WHERE user_id = $1 AND chat_id = $2 AND created >= $3`,
		userID,
		chatID,
		date,
	)
	if err != nil {
		return false, err
	}

	return rows.Next(), nil
}
