package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hellodoge/courses-tg-bot/courses"
	predefinedMessages "github.com/hellodoge/courses-tg-bot/courses/messages"
	"github.com/hellodoge/courses-tg-bot/internal/telegram/messages"
	"github.com/sirupsen/logrus"
)

const (
	authCommand         = "auth"
	newCommand          = "new"
	addModeratorCommand = "moder"
	getCourseCommand    = "course"
)

func (b *Bot) handleCommand(command *tgbotapi.Message) error {
	switch command.Command() {
	case "start":
		return b.handleCommandStart(command)
	case authCommand:
		return b.handleCommandAuth(command)
	case newCommand:
		return b.handleCommandNew(command)
	case addModeratorCommand:
		return b.handleCommandAddModerator(command)
	case getCourseCommand:
		return b.handleCommandGetCourse(command)
	default:
		unknownCommandMessage := tgbotapi.NewMessage(command.Chat.ID, predefinedMessages.UnknownCommand)
		_, err := b.bot.Send(unknownCommandMessage)
		return err
	}
}

func (b *Bot) handleCommandStart(command *tgbotapi.Message) error {
	err := b.service.Telegram.NewUser(command.Chat.ID)
	if err != nil {
		logrus.Error(err)
	}
	_, err = b.bot.Send(messages.PrepareStartMessage(command))
	return err
}

func (b *Bot) handleCommandAuth(command *tgbotapi.Message) error {
	err := b.service.Telegram.SetUserToken(command.Chat.ID, command.CommandArguments())
	if err != nil {
		return err
	}
	_, err = b.bot.Send(tgbotapi.NewMessage(command.Chat.ID, predefinedMessages.TokenSuccessfullySet))
	return err
}

func (b *Bot) handleCommandNew(command *tgbotapi.Message) error {
	isModerator, err := b.service.Telegram.UserIsModerator(command.Chat.ID)
	if err != nil {
		return err
	}

	if !isModerator {
		message := tgbotapi.NewMessage(command.Chat.ID, predefinedMessages.NotAModerator)
		_, err := b.bot.Send(message)
		return err
	}

	var course = courses.Course{
		Title: command.CommandArguments(),
	}

	id, err := b.service.Courses.NewCourse(&course)
	if err != nil {
		if err := b.ReplyWithText(command.Chat.ID, command.MessageID, err.Error()); err != nil {
			return err
		}
		return err
	}

	return b.ReplyWithText(command.Chat.ID, command.MessageID, id)
}

func (b *Bot) handleCommandAddModerator(command *tgbotapi.Message) error {
	isModerator, err := b.service.Telegram.UserIsAdmin(command.Chat.ID)
	if err != nil {
		return err
	}

	if !isModerator {
		return b.SendText(command.Chat.ID, predefinedMessages.NotAnAdministrator)
	}

	token, err := b.service.NewModerator(command.CommandArguments())
	if err != nil {
		return err
	}

	return b.ReplyWithText(command.Chat.ID, command.MessageID, token)
}

func (b *Bot) handleCommandGetCourse(command *tgbotapi.Message) error {
	course, err := b.service.GetCourse(command.CommandArguments())
	if err != nil {
		return err
	}
	if course == nil {
		return b.ReplyWithText(command.Chat.ID, command.MessageID, predefinedMessages.InvalidCourseID)
	}
	return b.sendCourseDescription(command.Chat.ID, course)
}
