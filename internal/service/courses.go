package service

import (
	"github.com/hellodoge/courses-tg-bot/courses"
	"github.com/hellodoge/courses-tg-bot/courses/messages"
	"github.com/hellodoge/courses-tg-bot/internal/repository"
	"strings"
)

type CoursesService struct {
	repo repository.Courses
}

func NewCoursesService(repo repository.Courses) *CoursesService {
	return &CoursesService{repo: repo}
}

func (s *CoursesService) NewCourse(course *courses.Course) (string, error) {
	course.Title = strings.TrimSpace(course.Title)
	if course.Title == "" {
		return "", Error{
			userError: messages.CourseTitleCannotBeEmpty,
		}
	}
	return s.repo.NewCourse(course)
}

func (s *CoursesService) GetCourse(id string) (*courses.Course, error) {
	course, err := s.repo.GetCourse(id)
	if err != nil {
		return nil, err
	} else if course == nil {
		return nil, Error{
			userError: messages.InvalidCourseID,
		}
	}
	return course, nil
}

func (s *CoursesService) GetLesson(id string) (*courses.Lesson, error) {
	lesson, err := s.repo.GetLesson(id)
	if err != nil {
		return nil, err
	} else if lesson == nil {
		return nil, Error{
			userError: messages.InvalidLessonID,
		}
	}
	return lesson, nil
}
