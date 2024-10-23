package main

import (
	"fmt"
	"main/internal/config"
	"net"
	"sync"

	"github.com/rs/zerolog/log"
)

type Server struct {
	listener net.Listener
	wg       sync.WaitGroup
	conf     config.Config
	router   *EventRouter
}

func NewServer(conf config.Config) *Server {
	s := &Server{
		conf: conf,
	}

	return s
}

func (s *Server) Init() {
	router := NewEventRouter()
	router.RegisterHandler("ECHO", &EchoHandler{})
	router.RegisterHandler("PING", &PingHandler{})

	s.router = router
}

func (s *Server) Run() error {
	var err error
	s.listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", s.conf.Service.Ip, s.conf.Service.Port))
	if err != nil {
		log.Err(err).Msg("Failed to Listen")
		return err
	}
	defer s.listener.Close()

	log.Info().Msgf("Sever running... %s", s.listener.Addr())

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Err(err).Msg("Failed to Accept")
			return err
		}

		go func(c net.Conn) {
			session := NewSession(c, s.conf.BufferSize)
			defer func() {
				session.Close()
				s.wg.Done()
			}()

			// socket Open
			session.Open()
			s.wg.Add(1)

			for {
				msg, err := session.Read()
				if err != nil {
					log.Err(err).Msg("Read error")
				}

				s.router.RouteMessage(session, msg)
			}
		}(conn)
	}
}

// Stop server graceful shutdown
func (s *Server) Stop() {
	log.Info().Msg("Server stopping...")

	err := s.listener.Close()
	if err != nil {
		log.Err(err).Msg("Failed to Close")
	}

	s.wg.Wait()

	log.Info().Msg("Server gracefully shutdown")
}
