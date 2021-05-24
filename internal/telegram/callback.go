package telegram

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hellodoge/courses-tg-bot/internal/telegram/callback"
	"github.com/sirupsen/logrus"
)

type callbackButton struct {
	text  string
	query callback.Query
}

func NewKeyboard(buttons ...callbackButton) (tgbotapi.InlineKeyboardMarkup, error) {
	var rows = make([][]tgbotapi.InlineKeyboardButton, 0, len(buttons))
	for _, button := range buttons {
		encoded, err := json.Marshal(button.query)
		if err != nil {
			return tgbotapi.InlineKeyboardMarkup{}, err
		}
		inlineButton := tgbotapi.NewInlineKeyboardButtonData(button.text, string(encoded))
		rows = append(rows, tgbotapi.NewInlineKeyboardRow(inlineButton))
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...), nil
}

func (b *Bot) handleCallback(chatID int64, callbackQuery *tgbotapi.CallbackQuery) error {
	var query callback.Query
	err := json.Unmarshal([]byte(callbackQuery.Data), &query)
	if err != nil {
		logrus.Info(err)
		return nil
	}
	switch query.Action {
	case callback.ActionGetLesson:
		return b.handleCallbackGetLesson(chatID, query.ID)
	case callback.ActionGetCourseLessons, callback.ActionGetCourseDescription:
		return fmt.Errorf("handleCallback: callback %s not implemented yet", query.Action)
	default:
		return nil
	}
}

func (b *Bot) handleCallbackGetLesson(chatID int64, lessonID string) error {
	lesson, err := b.service.GetLesson(lessonID)
	if err != nil {
		return err
	} else if lesson == nil {
		return nil
	}
	return b.sendLesson(chatID, lesson)
}
