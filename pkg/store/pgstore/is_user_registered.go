package pgstore

func (s *PGStore) IsUserRegistered(userID int, chatID int64) (bool, error) {
	rows, err := s.db.Query(
		`SELECT NULL FROM "user" WHERE user_id = $1 AND chat_id = $2`,
		userID,
		chatID,
	)
	if err != nil {
		return false, err
	}

	return rows.Next(), nil
}
