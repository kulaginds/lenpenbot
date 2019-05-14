package pgstore

func (s *PGStore) SetTop(chatID int64, top string) error {
	hasTop, err := s.HasTop(chatID)
	if err != nil {
		return err
	}

	if hasTop {
		return s.UpdateTop(chatID, top)
	}

	return s.CreateTop(chatID, top)
}
