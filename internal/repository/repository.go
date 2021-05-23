package repository

import (
	"github.com/hellodoge/courses-tg-bot/courses"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type Roles interface {
	NewAdmin(token, description string) error
	NewModerator(token, description string) error
}

type Courses interface {
	NewCourse(course *courses.Course) (string, error)
}

type Repository struct {
	Telegram *Telegram
	Courses
	Roles
}

func NewRepository(db *sqlx.DB, client *mongo.Client) *Repository {
	return &Repository{
		Telegram: NewTelegram(db),
		Courses:  NewCoursesMongoDB(client),
		Roles:    NewRolesMySQL(db),
	}
}
