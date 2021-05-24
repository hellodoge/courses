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

func (b *Bot) sendLesson(chatID int64, lesson *courses.Lesson) error {
	var (
		messageList     []tgbotapi.Chattable
		lastMessageBase *tgbotapi.BaseChat
	)
	if lesson.Title != "" || lesson.Description != "" {
		message := tgbotapi.NewMessage(chatID, fmt.Sprintf("%s\n%s", lesson.Title, lesson.Description))
		messageList = append(messageList, &message)
		lastMessageBase = &message.BaseChat
	}
	for _, document := range lesson.Documents {
		if shortcut, ok := document.URLs.Shortcuts[courses.TelegramShortcut]; ok {
			switch document.Type {
			case courses.TypeDocument:
				message := tgbotapi.NewDocumentShare(chatID, shortcut)
				message.Caption = document.Description
				messageList = append(messageList, &message)
				lastMessageBase = &message.BaseChat
			case courses.TypePhoto:
				message := tgbotapi.NewPhotoShare(chatID, shortcut)
				message.Caption = document.Description
				messageList = append(messageList, &message)
				lastMessageBase = &message.BaseChat
			case courses.TypeVideo:
				message := tgbotapi.NewVideoShare(chatID, shortcut)
				message.Caption = document.Description
				messageList = append(messageList, &message)
				lastMessageBase = &message.BaseChat
			default:
				logrus.Errorf("sendLesson: unknown document type: %s (lesson id: %s)", document.Type, lesson.ID)
				continue
			}
		} else {
			return errors.New("sendLesson: load files from url not implemented yet")
		}
	}
	if lesson.NextLessonID != "" {
		keyboard, err := NewKeyboard(callbackButton{
			text: messages.NextLesson,
			query: callback.Query{
				Action: callback.ActionGetLesson,
				ID:     lesson.NextLessonID,
			},
		})
		if err != nil {
			return err
		}
		if lastMessageBase == nil {
			message := tgbotapi.NewMessage(chatID, messages.LessonContentNotAddedYet)
			messageList = append(messageList, &message)
			lastMessageBase = &message.BaseChat
		}
		lastMessageBase.ReplyMarkup = keyboard
	}
	for _, message := range messageList {
		_, err := b.bot.Send(message)
		if err != nil {
			return err
		}
	}
	return nil
}
