package telegram

import (
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hellodoge/courses-tg-bot/courses"
	"github.com/hellodoge/courses-tg-bot/courses/messages"
	"github.com/hellodoge/courses-tg-bot/internal/telegram/callback"
)

func (b *Bot) SendCourseDescription(chatID int64, course *courses.Course) error {
	var (
		message tgbotapi.Chattable
		base    *tgbotapi.BaseChat
	)
	var description = fmt.Sprintf("%s\n\n%s", course.Title, course.Description)
	if course.Preview != nil {
		if shortcut, ok := course.Preview.URLs.Shortcuts[courses.TelegramShortcut]; ok {
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
