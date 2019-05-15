package pgstore

import "time"

func (s *PGStore) Enlarge(userID int, chatID int64, length int, created time.Time) error {
	created = created.Truncate(24 * time.Hour)
	_, err := s.db.Exec(`INSERT INTO "enlarge" (user_id, chat_id, length, created) VALUES ($1, $2, $3, $4)`, userID, chatID, length, created)

	return err
}
