package main

import (
	app "ValREST/internal/app"
	"log"
)

func main() {
	serv := new(app.Server)
	if err := serv.Start(); err != nil {
		log.Fatalf("error starting server: %s", err.Error())
	}
}
