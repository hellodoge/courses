package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleUpdate(update tgbotapi.Update) error {
	if update.Message == nil {
		return nil
	}

	if update.Message.Photo != nil {
		return b.handlePhoto(update.Message.Chat.ID, update.Message.MessageID, update.Message.Photo)
	}

	if update.Message.Video != nil {
		return b.handleVideo(update.Message.Chat.ID, update.Message.MessageID, update.Message.Video)
	}

	if update.Message.Document != nil {
		return b.handleDocument(update.Message.Chat.ID, update.Message.MessageID, update.Message.Document)
	}

	if update.Message.IsCommand() {
		return b.handleCommand(update.Message)
	}

	return b.handleMessage(update.Message)
}
