package main

import (
	"net"
	"sync"

	"github.com/rs/zerolog/log"
)

type Server struct {
	listener net.Listener
	wg       sync.WaitGroup
	// config
}

func NewServer() *Server {
	s := &Server{}

	return s
}

func (s *Server) Init() {
}

func (s *Server) Run() error {
	var err error
	s.listener, err = net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Err(err).Msg("Failed to Listen")
		return err
	}
	defer s.listener.Close()

	log.Info().Msg("Sever running...")

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Err(err).Msg("Failed to Accept")
			return err
		}

		go func(c net.Conn) {
			session := NewSession(c, 1024)
			defer func() {
				session.Close()
				s.wg.Done()
			}()

			// socket Open
			log.Info().Msg("Session opened")
			s.wg.Add(1)

			// socket Read/Write
			// TODO read / write

			s.wg.Wait()
		}(conn)
	}
}

// Stop server graceful shutdown
func (s *Server) Stop() {
	err := s.listener.Close()
	if err != nil {
		log.Err(err).Msg("Failed to Close")
	}
}
