package service

import (
	"github.com/hellodoge/courses-tg-bot/courses"
	"github.com/hellodoge/courses-tg-bot/courses/messages"
	"github.com/hellodoge/courses-tg-bot/internal/repository"
)

type TgAuthService struct {
	repo repository.AuthTelegram
}

func NewTgAuthService(repo repository.AuthTelegram) *TgAuthService {
	return &TgAuthService{
		repo:repo,
	}
}

func (s *TgAuthService) NewUser(chatID int64) error {
	return s.repo.CreateUser(chatID)
}

func (s *TgAuthService) SetUserToken(chatID int64, token string) error {
	if len(token) != courses.TokenLength {
		return Error{
			userError:  messages.InvalidLengthOfToken,
		}
	}
	return s.repo.SetUserToken(chatID, token)
}

func (s *TgAuthService) UserIsAdmin(chatID int64) (bool, error) {
	return s.repo.UserIsAdmin(chatID)
}

func (s *TgAuthService) UserIsModerator(chatID int64) (bool, error) {
	return s.repo.UserIsModerator(chatID)
}