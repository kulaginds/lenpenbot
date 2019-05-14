package pgstore

import (
	"time"

	"github.com/kulaginds/lenpenbot/pkg/store"
)

func (s *PGStore) GetTodayEnlarges(chatID int64, today time.Time) ([]*store.Enlarge, error) {
	enlarges := make([]*store.Enlarge, 0)

	rows, err := s.db.Query(
		`SELECT
					user_id, chat_id, length, created
				FROM "enlarge"
				WHERE
					chat_id = $1
					AND created >= $2
				ORDER BY length DESC`,
		chatID,
		today,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		enlarge := new(store.Enlarge)

		err := rows.Scan(&enlarge.UserID, &enlarge.ChatID, &enlarge.Length, &enlarge.Created)
		if err != nil {
			return nil, err
		}

		enlarges = append(enlarges, enlarge)
	}

	return enlarges, nil
}
