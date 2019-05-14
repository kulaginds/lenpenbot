package pgstore

import "time"

func (s *PGStore) SetToday(chatID int64, today time.Time, todayTop string) error {
	todayDate := today.Truncate(24 * time.Hour)
	hasToday, err := s.HasToday(chatID, todayDate, today.Add(23*time.Hour+59*time.Minute+59*time.Second))
	if err != nil {
		return err
	}

	if hasToday {
		return s.UpdateToday(chatID, today, todayTop)
	}

	return s.CreateToday(chatID, today, todayTop)
}
