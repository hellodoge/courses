package telegram

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hellodoge/courses-tg-bot/courses/messages"
	"github.com/hellodoge/courses-tg-bot/internal/telegram/callback"
	"github.com/sirupsen/logrus"
	"strings"
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
	case callback.ActionGetCourseLessons:
		return b.handleCallbackGetLessons(chatID, query.ID)
	case callback.ActionGetCourseDescription:
		return fmt.Errorf("handleCallback: callback %s not implemented yet", query.Action)
	case callback.ActionSearch:
		return b.handleCallbackSearch(chatID, query.ID)
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

func (b *Bot) handleCallbackGetLessons(chatID int64, courseID string) error {
	course, err := b.service.GetCourse(courseID)
	if err != nil {
		return err
	}
	message := tgbotapi.NewMessage(chatID, course.Title)
	var buttons = make([]callbackButton, 0, len(course.Lessons))
	for i, lesson := range course.Lessons {
		var title string
		if strings.TrimSpace(lesson.Title) == "" {
			title = fmt.Sprintf(messages.LessonNTemplate, i+1)
		} else {
			title = lesson.Title
		}
		var button = callbackButton{
			text: title,
			query: callback.Query{
				Action: callback.ActionGetLesson,
				ID:     lesson.ID,
			},
		}
		buttons = append(buttons, button)
	}
	message.ReplyMarkup, err = NewKeyboard(buttons...)
	if err != nil {
		return err
	}
	_, err = b.bot.Send(message)
	return err
}

func (b *Bot) handleCallbackSearch(chatID int64, searchID string) error {
	courseList, err := b.service.SearchCoursesBySearchID(searchID, b.config.SearchMaxResults)
	if err != nil {
		return err
	}
	return b.sendSearchResults(chatID, courseList, searchID)
}
