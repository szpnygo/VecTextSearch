package main

import (
	"log"

	"github.com/szpnygo/VecTextSearch/config"
	"github.com/szpnygo/VecTextSearch/server"
)

func main() {
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	server.StartServer(appConfig)
}
