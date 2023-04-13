package main

import (
	"ValREST/internal/app"
	repo "ValREST/internal/database"
	postgres "ValREST/internal/database/postgresql"
	"ValREST/internal/handlers"
	"ValREST/internal/services"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to init db: %s", err.Error())
	}

	repos := repo.NewRepository(db)
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
