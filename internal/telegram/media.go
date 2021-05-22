package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handlePhoto(chatID int64, messageID int, photos *[]tgbotapi.PhotoSize) error {
	ok, err := b.service.Telegram.UserIsModerator(chatID)
	if err != nil || !ok {
		return err
	}
	var message string
	for _, photo := range *photos {
		message += fmt.Sprintf("%dx%d url: %s\n", photo.Height, photo.Width, photo.FileID)
	}
	return b.ReplyWithText(chatID, messageID, message)
}

func (b *Bot) handleVideo(chatID int64, messageID int, video *tgbotapi.Video) error {
	ok, err := b.service.Telegram.UserIsModerator(chatID)
	if err != nil || !ok {
		return err
	}
	return b.ReplyWithText(chatID, messageID, video.FileID)
}

func (b *Bot) handleDocument(chatID int64, messageID int, document *tgbotapi.Document) error {
	ok, err := b.service.Telegram.UserIsModerator(chatID)
	if err != nil || !ok {
		return err
	}
	return b.ReplyWithText(chatID, messageID, document.FileID)
}
