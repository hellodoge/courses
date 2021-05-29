package telegram

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hellodoge/courses-tg-bot/courses"
	"github.com/hellodoge/courses-tg-bot/courses/messages"
	"github.com/hellodoge/courses-tg-bot/internal/telegram/callback"
	"github.com/sirupsen/logrus"
)

func (b *Bot) sendCourseDescription(chatID int64, course *courses.Course) error {
	var (
		message tgbotapi.Chattable
		base    *tgbotapi.BaseChat
	)
	var description = fmt.Sprintf("%s\n\n%s", course.Title, course.Description)
	if course.Preview != nil {
		if shortcut, ok := course.Preview.Shortcuts[courses.TelegramShortcut]; ok {
			msg := tgbotapi.NewPhotoShare(chatID, shortcut)
			msg.Caption = description
			base = &msg.BaseChat
			message = &msg
		} else {
			return errors.New("sendCourseDescription: loading photo into telegram not implemented yet")
		}
	} else {
		msg := tgbotapi.NewMessage(chatID, description)
		base = &msg.BaseChat
		message = &msg
	}

	if course.Lessons != nil && len(course.Lessons) > 0 {
		keyboard, err := NewKeyboard(
			callbackButton{
				text: messages.StartLearningCourse,
				query: callback.Query{
					Action: callback.ActionGetLesson,
					ID:     course.Lessons[0].ID,
				},
			},
			callbackButton{
				text: messages.GetCourseLessons,
				query: callback.Query{
					Action: callback.ActionGetCourseLessons,
					ID:     course.ID,
				},
			},
		)
		if err != nil {
			return err
		}
		base.ReplyMarkup = keyboard
	}
	_, err := b.bot.Send(message)
	return err
}

func (b *Bot) sendSearchResults(chatID int64, courseList []courses.Course, searchID string) error {
	var buttons []callbackButton
	for _, current := range courseList {
		if current.Title == "" {
			logrus.Errorf("Course %s title is empty, skipping", current.ID)
			continue
		}
		buttons = append(buttons, callbackButton{
			text: current.Title,
			query: callback.Query{
				Action: callback.ActionGetCourseDescription,
				ID:     current.ID,
			},
		})
	}
	if len(buttons) == 0 {
		return b.SendText(chatID, messages.NoMoreSearchResult)
	}
	if int64(len(courseList)) == b.config.SearchMaxResults && searchID != "" {
		buttons = append(buttons, callbackButton{
			text: messages.MoreResults,
			query: callback.Query{
				Action: callback.ActionSearch,
				ID:     searchID,
			},
		})
	}
	keyboard, err := NewKeyboard(buttons...)
	if err != nil {
		return err
	}
	message := tgbotapi.NewMessage(chatID, messages.SearchResults)
	message.ReplyMarkup = keyboard
	_, err = b.bot.Send(message)
	return err
}
