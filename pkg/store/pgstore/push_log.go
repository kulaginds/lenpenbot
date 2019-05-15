package pgstore

import (
	"errors"
	"fmt"
	"time"
)

const TextMaxLength = 10 * 1024 * 1024

func (s *PGStore) PushLog(text string) error {
	_, err := s.db.Exec(`
		INSERT INTO "log" (id, text, created)
		VALUES (NULL, $1, $2)`,
		text[:TextMaxLength],
		time.Now().UTC(),
	)

	if err != nil {
		err = errors.New(fmt.Sprintf("PushLog error: %s: %s", text, err))
	}

	return err
}