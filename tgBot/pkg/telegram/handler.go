package telegram

import (
	"context"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error{
	switch message.Command(){
	case startCommand:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	if _, err := b.getAccessToken(message); err != nil {
		return b.initAuth(message)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.AlreadyAuthorized)
			
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.UnknownCommand)
			
	_, err := b.bot.Send(msg)

	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, b.messages.SavedSuccessfully)
			
	accessToken, err := b.getAccessToken(message)
	if err != nil {
		return errorAuth
	}

	_, err = url.ParseRequestURI(message.Text)
	if err != nil {
		return errorUrl
	}

	err = b.pocketClient.Add(context.Background(), pocket.AddInput{AccessToken: accessToken, URL: message.Text})
	if err != nil {
		return errorSave
	}

	_, err = b.bot.Send(msg)

	return err
}