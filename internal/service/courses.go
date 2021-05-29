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

func (s *CoursesService) SearchCourses(query string, limit, offset int64) ([]courses.Course, error) {
	if query == "" {
		return nil, Error{
			userError: messages.SearchQueryCannotBeEmpty,
		}
	}
	return s.repo.SearchCourses(query, limit, offset)
}

func (s *CoursesService) SearchCoursesBySearchID(searchID string, limit int64) ([]courses.Course, error) {
	result, err := s.repo.SearchCoursesBySearchID(searchID, limit)
	if err == repository.ErrorInvalidSearchID {
		return nil, Error{
			userError: messages.InvalidSearchID,
		}
	}
	return result, err
}

func (s *CoursesService) NewSearch(query string, offset int64) (string, error) {
	if query == "" {
		return "", Error{
			userError: messages.SearchQueryCannotBeEmpty,
		}
	}
	return s.repo.NewSearch(query, offset)
}
