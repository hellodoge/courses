package service

import "github.com/hellodoge/courses-tg-bot/internal/repository"

type TgAuthorization interface {
	NewUser(chatID int64) error
	SetUserToken(chatID int64, token string) error
	UserIsAdmin(chatID int64) (bool, error)
	UserIsModerator(chatID int64) (bool, error)
}

type TelegramService struct {
	TgAuthorization
}

func NewTelegramService(repo *repository.Telegram) *TelegramService {
	return &TelegramService{
		TgAuthorization: NewTgAuthService(repo.AuthTelegram),
	}
}