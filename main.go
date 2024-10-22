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

	s := NewServer(conf)
	defer s.Stop()

	s.Init()

	if err := s.Run(); err != nil {
		log.Err(err).Msg("Server failed to run")
	}
}
