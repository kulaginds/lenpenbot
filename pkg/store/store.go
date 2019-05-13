package store

type Store interface {
	IsUserRegistered(userID int, chatID int64) (bool, error)
	RegisterUser(userID int, chatID int64) error
}
