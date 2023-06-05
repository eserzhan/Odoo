package main

import (
	"context"
	"log"

	shop "github.com/eserzhan/clotheStore"
	"github.com/eserzhan/clotheStore/pkg/handler"
	"github.com/eserzhan/clotheStore/pkg/repository"
	"github.com/eserzhan/clotheStore/pkg/service"
	//"github.com/spf13/viper"
)

func main() {
	//err := initConfig()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	mongoClient, err := repository.NewMongoDB("admin", "qwerty")
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(context.Background())

	_ = mongoClient.Database("mongodb")

	service := service.NewService()
	handler := handler.NewHandler(service)

	srv := new(shop.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

// func initConfig() error {
// 	viper.AddConfigPath("configs")
// 	viper.SetConfigName("config")
// 	return viper.ReadInConfig()
// }