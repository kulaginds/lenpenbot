package store

import "time"

type User struct {
	UserID int32 `db:"id"`
	ChatID int32 `db:"chat_id"`
}

type Enlarge struct {
	UserID  int32     `db:"id"`
	ChatID  int32     `db:"chat_id"`
	Length  int       `db:"length"`
	Created time.Time `db:"created"`
}

type TopType string

const (
	TopTypeAll   TopType = "all"
	TopTypeToday TopType = "today"
)

type Top struct {
	ChatID  int32   `db:"chat_id"`
	Type    TopType `db:"type"`
	Message string  `db:"message"`
}

type Credit struct {
	UserID   int32     `db:"user_id"`
	ChatID   int32     `db:"chat_id"`
	Length   int       `db:"length"`
	Percents int       `db:"percents"`
	Repaid   bool      `db:"repaid"`
	Created  time.Time `db:"created"`
}
