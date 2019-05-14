package store

import "time"

type Store interface {
	IsUserRegistered(userID int, chatID int64) (bool, error)
	RegisterUser(userID int, chatID int64) error

	IsEnlarge(userID int, chatID int64, date time.Time) (bool, error)
	Enlarge(userID int, chatID int64, length int, created time.Time) error
}
