package main

import (
	"log"
	"net/http"
	"zvon/config"
	"zvon/logger"
	"zvon/router"
	"zvon/state"
)

func main() {
	// Load config and create shared state.
	config.Load()
	state.CreateState()

	// Set minimum logger level from config.
	logger.MinimumLevel = logger.Level(config.Current.MinimumLogLevel)

	go Listen()

	// Create router and run server.
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(config.Current.ListenAddress, r))
}
