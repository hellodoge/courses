package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	_, err := b.bot.Send(msg)
	return err
}