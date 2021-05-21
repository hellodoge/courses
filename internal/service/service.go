package service

import "github.com/hellodoge/courses-tg-bot/internal/repository"

type Service struct {
	Telegram *TelegramService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Telegram: NewTelegramService(repository.Telegram),
	}
}