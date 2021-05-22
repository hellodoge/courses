package service

import "github.com/hellodoge/courses-tg-bot/internal/repository"

type Roles interface {
	NewAdmin(description string) (string, error)
	NewModerator(description string) (string, error)
}

type Service struct {
	Telegram *TelegramService
	Roles
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Telegram: NewTelegramService(repository.Telegram),
		Roles:    NewRolesService(repository.Roles),
	}
}
