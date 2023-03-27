package main

import (
	"log"

	"github.com/szpnygo/VecTextSearch/api"
	"github.com/szpnygo/VecTextSearch/config"
)

func main() {
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	api.StartServer(appConfig)
}
