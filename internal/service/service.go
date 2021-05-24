package service

import (
	"github.com/hellodoge/courses-tg-bot/courses"
	"github.com/hellodoge/courses-tg-bot/internal/repository"
)

type Roles interface {
	NewAdmin(description string) (string, error)
	NewModerator(description string) (string, error)
	UserIsAdmin(token string) (bool, error)
	UserIsModerator(token string) (bool, error)
}

type Courses interface {
	NewCourse(course *courses.Course) (string, error)
	GetCourse(id string) (*courses.Course, error)
	GetLesson(id string) (*courses.Lesson, error)
}

type Service struct {
	Telegram *TelegramService
	Courses
	Roles
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Telegram: NewTelegramService(repository.Telegram),
		Courses:  NewCoursesService(repository.Courses),
		Roles:    NewRolesService(repository.Roles),
	}
}
