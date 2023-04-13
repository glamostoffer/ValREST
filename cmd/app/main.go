package main

import (
	"ValREST/internal/app"
	"ValREST/internal/handlers"
	"log"
)

func main() {
	serv := new(app.Server)
	handler := new(handlers.Handler)
	if err := serv.Start(handler.InitRouts()); err != nil {
		log.Fatalf("error starting server: %s", err.Error())
	}
}
