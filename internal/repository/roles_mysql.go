package repository

import "github.com/jmoiron/sqlx"

const (
	mySqlCreateAdmin     = "create_admin_token.sql"
	mySqlCreateModerator = "create_moderator_token.sql"
)

type RolesMySQL struct {
	db *sqlx.DB
}

func NewRolesMySQL(db *sqlx.DB) *RolesMySQL {
	return &RolesMySQL{db: db}
}

func (r *RolesMySQL) NewAdmin(token, description string) error {
	createAdminQuery, err := getQuery(mySqlQueriesFolder, mySqlCreateAdmin)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(createAdminQuery, token, description)
	return err
}

func (r *RolesMySQL) NewModerator(token, description string) error {
	createModeratorQuery, err := getQuery(mySqlQueriesFolder, mySqlCreateModerator)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(createModeratorQuery, token, description)
	return err
}
