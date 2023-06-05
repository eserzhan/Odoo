package main

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/eserzhan/tgBott/pkg/config"
	"github.com/eserzhan/tgBott/pkg/repository"
	"github.com/eserzhan/tgBott/pkg/repository/boltDb"
	"github.com/eserzhan/tgBott/pkg/server"
	"github.com/eserzhan/tgBott/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

const (
	//redirectURL = "http://localhost/"
	//ru = "https://telegram.org/#6004117074"
	//ru = "https://google.com"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(cfg)
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}
	
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	pocket, err := pocket.NewClient(cfg.PocketConsumerKey)

	if err != nil {
		log.Fatal(err)
	}

	db, err := initBolt(cfg)
	if err != nil {
		log.Fatal(err)
	}
	storage := boltDb.NewBolt(db)

	server := server.NewServer(pocket, storage, cfg.TelegramBotURL)

	tgBot := telegram.NewBot(bot, pocket, storage, cfg.AuthServerURL, cfg.Messages)
	

	go func() {
		tgBot.Run()
		
	}()

	if err = server.Start(); err != nil {
		log.Fatal(err)
	}
}

func initBolt(cfg *config.Config) (*bolt.DB, error) {
	db, err := bolt.Open(cfg.DBpath, 0600, nil)
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccessToken))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestToken))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return db, nil
}
