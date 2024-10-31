package main

import (
	"gotcp/internal/bs"

	"github.com/rs/zerolog/log"
)

type EventHandler interface {
	Handle(session *Session, message string)
}

// EchoHandler 에코 서버
type EchoHandler struct{}

func (e *EchoHandler) Handle(session *Session, message string) {
	log.Info().Msgf("Echoing message: %s", message)
	session.Write(bs.StringToBytes(message))
}

// PingHandler PING 서버
type PingHandler struct{}

func (p *PingHandler) Handle(session *Session, message string) {
	log.Info().Msg("Received PING, responding with PONG")
	session.Write(bs.StringToBytes("PONG"))
}
