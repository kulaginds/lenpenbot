package pgstore

import (
	"errors"
	"fmt"
	"time"

	"github.com/kulaginds/lenpenbot/pkg/store"
)

func (s *PGStore) GetToday(chatID int64, today time.Time) (*store.Top, error) {
	top := new(store.Top)

	rows, err := s.db.Query(
		`SELECT
					chat_id, type, message, updated, created
				FROM "top"
				WHERE
					chat_id = $1
				  	AND type = $2
					AND updated >= $3`,
		chatID,
		store.TopTypeToday,
		today,
	)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, errors.New(fmt.Sprintf("GetToday: no today top: chatID=%d; updated=%s", chatID, today))
	}

	err = rows.Scan(&top.ChatID, &top.Type, &top.Message, &top.Updated, &top.Created)
	if err != nil {
		return nil, err
	}

	return top, nil
}
