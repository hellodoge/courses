package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hellodoge/courses-tg-bot/pkg/util"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	courses, err := b.service.SearchCourses(message.Text)
	if err != nil {
		return err
	}
	var search = ""
	if int64(len(courses)) > b.config.SearchMaxResults {
		search, err = b.service.NewSearch(message.Text, courses, int64(len(courses)))
		if err != nil {
			return err
		}
	}
	to := util.MinInt64(b.config.SearchMaxResults, int64(len(courses)))
	return b.sendSearchResults(message.Chat.ID, courses[:to], search)
}
