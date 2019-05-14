package pgstore

import "github.com/kulaginds/lenpenbot/pkg/store"

func (s *PGStore) GetTopEnlarges(chatID int64, limit int) ([]*store.Enlarge, error) {
	enlarges := make([]*store.Enlarge, 0)

	rows, err := s.db.Query(
		`SELECT
    				DISTINCT user_id,
                    MAX(length) AS l
				FROM "enlarge"
				WHERE
				    chat_id = $1
				GROUP BY user_id
				ORDER BY l DESC
				limit $2`,
		chatID,
		limit,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		enlarge := new(store.Enlarge)

		err := rows.Scan(&enlarge.UserID, &enlarge.Length)
		if err != nil {
			return nil, err
		}

		enlarges = append(enlarges, enlarge)
	}

	return enlarges, nil
}
