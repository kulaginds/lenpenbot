package store

import "time"

type Store interface {
	IsUserRegistered(userID int, chatID int64) (bool, error)
	RegisterUser(userID int, chatID int64) error

	IsEnlarge(userID int, chatID int64, date time.Time) (bool, error)
	Enlarge(userID int, chatID int64, length int, created time.Time) error

	HasToday(chatID int64, updatedMin, updatedMax time.Time) (bool, error)
	HasTop(chatID int64) (bool, error)
	GetToday(chatID int64, today time.Time) (*Top, error)
	GetTop(chatID int64) (*Top, error)
	UpdateToday(chatID int64, today time.Time) error
	UpdateTop(chatID int64) error
}
