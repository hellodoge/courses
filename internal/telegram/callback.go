package telegram

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hellodoge/courses-tg-bot/internal/telegram/callback"
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
