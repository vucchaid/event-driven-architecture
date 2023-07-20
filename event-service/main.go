package main

import (
	"log"

	"github.com/vucchaid/event-driven-architecture/event-service/api"
	"github.com/vucchaid/event-driven-architecture/event-service/config"
)

func main() {
	config, err := config.GetApplicationConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := api.Serve(config.Port, config.Connection, config.DBtype); err != nil {
		log.Fatal(err)
	}
}
