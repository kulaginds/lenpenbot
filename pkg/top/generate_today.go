package top

import "time"

func (t *Top) GenerateToday(chatID int64, today time.Time) (string, error) {
	enlarges, err := t.store.GetTodayEnlarges(chatID, today)
	if err != nil {
		return "", err
	}

	nicksMap, err := t.getNicksMap(t.extractUserIDsFromEnlarges(enlarges), chatID)
	if err != nil {
		return "", err
	}

	return t.generateTop(enlarges, nicksMap), nil
}
