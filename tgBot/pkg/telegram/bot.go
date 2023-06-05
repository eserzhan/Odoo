package telegram

import (
	"github.com/eserzhan/tgBott/pkg/config"
	"github.com/eserzhan/tgBott/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

const (
	startCommand = "start"

	//startMessage = "Привет! Чтобы начать сохранять ссылки в своем Pocket аккаунте, для начала тебе необходимо дать мне на это доступ. Для этого переходи по ссылке:\n%s"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	pocketClient *pocket.Client
	db repository.TokenRepository
	redirectURL string
	messages config.Messages
}

func NewBot(bot *tgbotapi.BotAPI, pocketClient *pocket.Client, db repository.TokenRepository, redirectURL string, messages config.Messages) *Bot {
	return &Bot{bot: bot, pocketClient: pocketClient, db: db, redirectURL: redirectURL, messages: messages}
}

func (b *Bot) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue 
		}

		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.errorHandler(update.Message.Chat.ID, err)
			}
			continue
		}

		if err := b.handleMessage(update.Message); err != nil {
			b.errorHandler(update.Message.Chat.ID, err)
		}
	}

	
}

