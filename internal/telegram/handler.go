package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleUpdate(update tgbotapi.Update) error {
	if update.Message == nil {
		return nil
	}

	if update.Message.IsCommand() {
		return b.handleCommand(update.Message)
	}

	return b.handleMessage(update.Message)
}

