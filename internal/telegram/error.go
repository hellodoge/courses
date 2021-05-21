package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hellodoge/courses-tg-bot/internal/service"
	"github.com/sirupsen/logrus"
)

func (b *Bot) HandleServiceErrors(serviceErr service.Error, update tgbotapi.Update) {

	if serviceErr.IsSystemError() {
		logrus.Error(serviceErr.Log())
	}

	if update.Message != nil && serviceErr.HasUserError() {
		message := tgbotapi.NewMessage(update.Message.Chat.ID, serviceErr.UserError())
		_, err := b.bot.Send(message)
		if err != nil {
			logrus.Error(err)
		}
	}
}