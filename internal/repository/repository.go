package repository

import "github.com/jmoiron/sqlx"

type Roles interface {
	NewAdmin(token, description string) error
	NewModerator(token, description string) error
}

type Repository struct {
	Telegram *Telegram
	Roles
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Telegram: NewTelegram(db),
		Roles:    NewRolesMySQL(db),
	}
}
