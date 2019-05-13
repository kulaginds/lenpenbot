package pgstore

import (
	"database/sql"
	"github.com/kulaginds/lenpenbot/pkg/store"
)

type PGStore struct {
	db *sql.DB
}

func NewPGStore(db *sql.DB) store.Store {
	return &PGStore{db: db}
}
