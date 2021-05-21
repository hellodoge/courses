package repository

import "github.com/jmoiron/sqlx"

type AuthTelegram interface {
	CreateUser(chatID int64) error
	SetUserToken(chatID int64, token string) error
	UserIsAdmin(chatID int64) (bool, error)
	UserIsModerator(chatID int64) (bool, error)
}

type Telegram struct {
	AuthTelegram
}

func NewTelegram(db *sqlx.DB) *Telegram {
	return &Telegram{
		AuthTelegram: NewTgAuthMySQL(db),
	}
}