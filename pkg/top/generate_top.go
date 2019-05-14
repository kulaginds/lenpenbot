package top

const topLimit = 10

func (t *Top) GenerateTop(chatID int64) (string, error) {
	enlarges, err := t.store.GetTopEnlarges(chatID, topLimit)
	if err != nil {
		return "", err
	}

	nicksMap, err := t.getNicksMap(t.extractUserIDsFromEnlarges(enlarges), chatID)
	if err != nil {
		return "", err
	}

	return t.generateTop(enlarges, nicksMap), nil
}
