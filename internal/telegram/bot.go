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
	SearchMaxResults     int64
	NumberOfWorkers      int
	MaxQueuedConnections int
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
	var updatesQueue = make(chan tgbotapi.Update, b.config.MaxQueuedConnections)
	for i := 0; i < b.config.NumberOfWorkers; i++ {
		go b.worker(updatesQueue)
	}
	for update := range updates {
		updatesQueue <- update
	}
}

func (b *Bot) worker(queue <-chan tgbotapi.Update) {
	for update := range queue {
		err := b.handleUpdate(update)
		if serviceErr, ok := err.(service.Error); ok {
			b.HandleServiceErrors(serviceErr, update)
		} else if err != nil {
			logrus.Error(err)
		}
	}
}
