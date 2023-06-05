package telegram

import (
	"context"
	"fmt"

	"github.com/eserzhan/tgBott/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
func (b *Bot) initAuth(message *tgbotapi.Message) error{
	authLink, err := b.generateAuthorizationLink(message.Chat.ID)

	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(b.messages.Start, authLink))
			
	_, err = b.bot.Send(msg)

	return err
}

func (b *Bot) getAccessToken(message *tgbotapi.Message) (string, error) {
	return b.db.Get(message.Chat.ID, repository.AccessToken)
}

func (b *Bot) generateAuthorizationLink(chatId int64) (string, error) {
	redirectUrl := b.generateRedirectURL(chatId)

	requestToken, err := b.pocketClient.GetRequestToken(context.Background(), redirectUrl)
	if err != nil {
		return "", err
	}

	if err = b.db.Save(chatId, requestToken, repository.RequestToken); err != nil {
		return "", err
	}

	return b.pocketClient.GetAuthorizationURL(requestToken, redirectUrl)
}

func (b *Bot) generateRedirectURL(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", b.redirectURL, chatID)
}