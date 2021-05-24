package service

import (
	"github.com/hellodoge/courses-tg-bot/courses"
	"github.com/hellodoge/courses-tg-bot/internal/repository"
)

type CoursesService struct {
	repo repository.Courses
}

func NewCoursesService(repo repository.Courses) *CoursesService {
	return &CoursesService{repo: repo}
}

func (s *CoursesService) NewCourse(course *courses.Course) (string, error) {
	return s.repo.NewCourse(course)
}

func (s *CoursesService) GetCourse(id string) (*courses.Course, error) {
	return s.repo.GetCourse(id)
}
