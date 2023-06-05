package telegram

import (
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	errorAuth = errors.New("user is unauthorized")

	errorUrl = errors.New("incorrect url address")

	errorSave = errors.New("unable to save")
)

func(b *Bot) errorHandler(chat_id int64, err error) {
	msg := tgbotapi.NewMessage(chat_id, b.messages.Default)
	
	switch err{
	case errorUrl:
		msg.Text = b.messages.InvalidUrl
		b.bot.Send(msg)
	case errorAuth:
		msg.Text = b.messages.Unauthorized
		b.bot.Send(msg)
	case errorSave:
		msg.Text = b.messages.UnableToSave
		b.bot.Send(msg)
	default:
		b.bot.Send(msg)
	}
}