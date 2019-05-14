package store

import "time"

type User struct {
	UserID   int `db:"id"`
	ChatID64 int `db:"chat_id"`
}

type Enlarge struct {
	UserID  int       `db:"id"`
	ChatID  int64     `db:"chat_id"`
	Length  int       `db:"length"`
	Created time.Time `db:"created"`
}

type TopType string

const (
	TopTypeAll   TopType = "all"
	TopTypeToday TopType = "today"
)

type Top struct {
	ChatID  int64     `db:"chat_id"`
	Type    TopType   `db:"type"`
	Message string    `db:"message"`
	Updated time.Time `db:"updated"`
	Created time.Time `db:"created"`
}

type Credit struct {
	UserID   int       `db:"user_id"`
	ChatID   int64     `db:"chat_id"`
	Length   int       `db:"length"`
	Percents int       `db:"percents"`
	Repaid   bool      `db:"repaid"`
	Created  time.Time `db:"created"`
}
