package main

import (
	"ValREST/internal/app"
	repo "ValREST/internal/database/postgrsql"
	"ValREST/internal/handlers"
	"ValREST/internal/services"
	"log"
)

func main() {
	repos := repo.NewRepository()
	services := services.NewService(repos)
	handler := handlers.NewHandler(services)

	serv := new(app.Server)
	if err := serv.Start(handler.InitRouts()); err != nil {
		log.Fatalf("error starting server: %s", err.Error())
	}
}
