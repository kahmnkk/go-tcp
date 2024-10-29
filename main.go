package main

import (
	"main/internal/config"
	"main/internal/logger"
	"os"

	"github.com/rs/zerolog/log"
)

func main() {
	// Config
	conf := config.Read()

	// Logger
	log.Logger = logger.NewLogger(conf.LogLevel, os.Stdout)

	server = NewServer(conf)
	defer server.Stop()

	server.Init()

	if err := server.Run(); err != nil {
		log.Err(err).Msg("Server failed to run")
	}
}
