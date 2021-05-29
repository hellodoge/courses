package repository

import (
	"github.com/hellodoge/courses-tg-bot/courses"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type Roles interface {
	NewAdmin(token, description string) error
	NewModerator(token, description string) error
	UserIsAdmin(token string) (bool, error)
	UserIsModerator(token string) (bool, error)
}

type Courses interface {
	SearchCourses(query string, limit, skip int64) ([]courses.Course, error)
	SearchCoursesBySearchID(searchID string, limit int64) ([]courses.Course, error)
	NewSearch(query string, skip int64) (string, error)
	NewCourse(course *courses.Course) (string, error)
	GetCourse(id string) (*courses.Course, error)
	GetLesson(idHex string) (*courses.Lesson, error)
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
