package repository

import "github.com/jmoiron/sqlx"


type Repository struct {
	Telegram *Telegram
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Telegram: NewTelegram(db),
	}
}