package main

import (
	"event-service/api"
	"event-service/config"
	"log"
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
