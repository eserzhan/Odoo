package main

import (
	"log"
	"os"

	"github.com/eserzhan/rest"
	"github.com/eserzhan/rest/pkg/handler"
	"github.com/eserzhan/rest/pkg/repository"
	"github.com/eserzhan/rest/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil{
		log.Fatalf("Ошибка чтения конфигурационного файла: %s", err)
	}

	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Ошибка чтения конфигурационного файла: %s", err)
    }

	db, err := repository.NewPostgresDB(repository.Configs{
		Port: viper.GetString("db.port"),
		Host: viper.GetString("db.host"),
		Dbname: viper.GetString("db.dbname"), 
		Sslmode: viper.GetString("db.sslmode"), 
		Username: viper.GetString("db.username"), 
		Password: os.Getenv("password")})
	
	if err != nil {
	log.Fatalf("can't to connect to postgres: %s", err)

	}
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)


	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil{
		log.Fatalf("Can't initialize server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}