package main

import (
	"main/internal/bs"
	"strings"

	"github.com/rs/zerolog/log"
)

type EventRouter struct {
	handlers map[string]EventHandler
}

func NewEventRouter() *EventRouter {
	return &EventRouter{
		handlers: make(map[string]EventHandler),
	}
}

func (r *EventRouter) RegisterHandler(command string, handler EventHandler) {
	r.handlers[command] = handler
}

// RouteMessage 패킷을 적절한 handler로 라우팅
func (r *EventRouter) RouteMessage(session *Session, message string) {
	// TODO protobuf
	command := strings.Split(message, " ")[0]

	if handler, exists := r.handlers[command]; exists {
		handler.Handle(session, message)
	} else {
		log.Warn().Msgf("Unknown command: %s", command)
		session.Write(bs.StringToBytes("ERROR: Unknown command\n"))
	}
}
