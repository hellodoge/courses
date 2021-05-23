package service

import (
	"github.com/hellodoge/courses-tg-bot/internal/repository"
	"github.com/hellodoge/courses-tg-bot/pkg/token"
)

type RolesService struct {
	repo repository.Roles
}

func NewRolesService(repo repository.Roles) *RolesService {
	return &RolesService{repo: repo}
}

func (r *RolesService) NewAdmin(description string) (string, error) {
	tok, err := token.GenerateToken()
	if err != nil {
		return "", err
	}
	err = r.repo.NewAdmin(tok, description)
	return tok, err
}

func (r *RolesService) NewModerator(description string) (string, error) {
	tok, err := token.GenerateToken()
	if err != nil {
		return "", err
	}
	err = r.repo.NewModerator(tok, description)
	return tok, err
}

func (r *RolesService) UserIsAdmin(token string) (bool, error) {
	return r.repo.UserIsAdmin(token)
}

func (r *RolesService) UserIsModerator(token string) (bool, error) {
	return r.repo.UserIsModerator(token)
}
