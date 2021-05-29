package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	courses, err := b.service.SearchCourses(message.Text, b.config.SearchMaxResults, 0)
	if err != nil {
		return err
	}
	var search = ""
	if int64(len(courses)) == b.config.SearchMaxResults {
		search, err = b.service.NewSearch(message.Text, int64(len(courses)))
		if err != nil {
			return err
		}
	}
	return b.sendSearchResults(message.Chat.ID, courses, search)
}
