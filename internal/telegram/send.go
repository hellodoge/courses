package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *Bot) SendText(chatID int64, message string) error {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) ReplyWithText(chatID int64, messageID int, message string) error {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ReplyToMessageID = messageID
	_, err := b.bot.Send(msg)
	return err
}
