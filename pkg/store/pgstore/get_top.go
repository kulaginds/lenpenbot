package pgstore

import (
	"errors"
	"fmt"

	"github.com/kulaginds/lenpenbot/pkg/store"
)

func (s *PGStore) GetTop(chatID int64) (*store.Top, error) {
	top := new(store.Top)

	rows, err := s.db.Query(
		`SELECT
					chat_id, type, message, updated, created
				FROM "top"
				WHERE
					chat_id = $1
					AND type = $2`,
		chatID,
		store.TopTypeAll,
	)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, errors.New(fmt.Sprintf("GetTop: no top: chatID=%d", chatID))
	}

	err = rows.Scan(&top.ChatID, &top.Type, &top.Message, &top.Updated, &top.Created)
	if err != nil {
		return nil, err
	}

	return top, nil
}
