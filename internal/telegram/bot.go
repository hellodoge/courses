package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hellodoge/courses-tg-bot/internal/service"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	service *service.Service
	config  Config
}

type Config struct {
	SearchMaxResults int64
}

func NewBot(bot *tgbotapi.BotAPI, service *service.Service, config Config) *Bot {
	return &Bot{
		bot:     bot,
		service: service,
		config:  config,
	}
}

func (b *Bot) InitUpdateChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) HandleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		err := b.handleUpdate(update)
		if serviceErr, ok := err.(service.Error); ok {
			b.HandleServiceErrors(serviceErr, update)
		} else if err != nil {
			logrus.Error(err)
		}
	}
}
