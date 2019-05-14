package pgstore

func (s *PGStore) RegisterUser(userID int, chatID int64) error {
	_, err := s.db.Exec(`INSERT INTO "user" (user_id, chat_id) VALUES ($1, $2)`, userID, chatID)

	return err
}
