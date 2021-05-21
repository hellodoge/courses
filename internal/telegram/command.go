package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	predefinedMessages "github.com/hellodoge/courses-tg-bot/courses/messages"
	"github.com/hellodoge/courses-tg-bot/internal/telegram/messages"
	"github.com/sirupsen/logrus"
)

const (
	authCommand = "auth"
	newCommand  = "new"
)

func (b *Bot) handleCommand(command *tgbotapi.Message) error {
	switch command.Command() {
	case "start":
		return b.handleCommandStart(command)
	case authCommand:
		return b.handleCommandAuth(command)
	case newCommand:
		return b.handleCommandNew(command)
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

	message := tgbotapi.NewMessage(command.Chat.ID, command.CommandArguments()) //TODO
	_, err = b.bot.Send(message)
	return err
}
