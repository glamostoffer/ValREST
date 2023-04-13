package main

import (
	"ValREST/internal/app"
	repo "ValREST/internal/database/postgrsql"
	"ValREST/internal/handlers"
	"ValREST/internal/services"
	"log"

	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repo.NewRepository()
	services := services.NewService(repos)
	handler := handlers.NewHandler(services)

	serv := new(app.Server)
	if err := serv.Start(viper.GetString("port"), handler.InitRouts()); err != nil {
		log.Fatalf("error starting server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
