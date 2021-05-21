package repository

import (
	"github.com/jmoiron/sqlx"
)

const (
	mySqlTgQueryCreateUser = "tg_create_user.sql"
	mySqlTgSetUserToken    = "tg_set_user_token.sql"
	mySqlTgUserIsAdmin     = "tg_user_is_admin.sql"
	mySqlTgUserIsModerator = "tg_user_is_moderator.sql"
)

type TgAuthMySQL struct {
	db *sqlx.DB
}

func NewTgAuthMySQL(db *sqlx.DB) *TgAuthMySQL {
	return &TgAuthMySQL{
		db: db,
	}
}

func (r *TgAuthMySQL) CreateUser(chatID int64) error {
	createUserQuery, err := getQuery(mySqlQueriesFolder, mySqlTgQueryCreateUser)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(createUserQuery, chatID)
	return err
}

func (r *TgAuthMySQL) SetUserToken(chatID int64, token string) error {
	setTokenQuery, err := getQuery(mySqlQueriesFolder, mySqlTgSetUserToken)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(setTokenQuery, token, chatID)
	return err
}

func (r *TgAuthMySQL) UserIsAdmin(chatID int64) (bool, error) {
	userIsAdminQuery, err := getQuery(mySqlQueriesFolder, mySqlTgUserIsAdmin)
	if err != nil {
		return false, err
	}
	var result bool
	err = r.db.Get(&result, userIsAdminQuery, chatID)
	return result, err
}

func (r *TgAuthMySQL) UserIsModerator(chatID int64) (bool, error) {
	userIsModeratorQuery, err := getQuery(mySqlQueriesFolder, mySqlTgUserIsModerator)
	if err != nil {
		return false, err
	}
	var result bool
	err = r.db.Get(&result, userIsModeratorQuery, chatID)
	return result, err
}