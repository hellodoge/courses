package messages

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hellodoge/courses-tg-bot/courses/messages"
)

func PrepareStartMessage(m *tgbotapi.Message) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(m.Chat.ID, messages.StartMessage)
}