package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken     string
	PocketConsumerKey string
	AuthServerURL     string
	TelegramBotURL    string `mapstructure:"bot_url"`
	DBpath            string `mapstructure:"db_file"`

	Messages Messages
}

type Messages struct {
	Errors
	Responses
}

type Errors struct {
	Default      string `mapstructure:"default"`
	InvalidUrl   string `mapstructure:"invalid_url"`
	Unauthorized string `mapstructure:"unauthorized"`
	UnableToSave string `mapstructure:"unable_to_save"`
}

type Responses struct {
	Start             string `mapstructure:"start"`
	AlreadyAuthorized string `mapstructure:"alreadyAuthorized"`
	UnknownCommand    string `mapstructure:"unknown_command"`
	SavedSuccessfully string `mapstructure:"saved_successfully"`
}

func InitConfig() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err 
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err 
	}

	if err := viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses); err != nil {
		return nil, err 
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err 
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err 
	}

	return &cfg, nil 
}

func parseEnv(cfg *Config) error {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	if err := viper.BindEnv("TGTOKEN"); err != nil {
		return err 
	}

	if err := viper.BindEnv("CONSUMER_KEY"); err != nil {
		return err 
	}

	if err := viper.BindEnv("AUTHSERVERURL"); err != nil {
		return err 
	}

	cfg.AuthServerURL = viper.GetString("AUTHSERVERURL")
	cfg.PocketConsumerKey = viper.GetString("CONSUMER_KEY")
	cfg.TelegramToken = viper.GetString("TGTOKEN")

	return nil 
}