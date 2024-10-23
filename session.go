package main

import (
	"io"
	"main/internal/bs"
	"net"

	"github.com/rs/zerolog/log"
)

type Session struct {
	conn net.Conn
	buf  []byte
}

func NewSession(conn net.Conn, bufSize int) *Session {
	s := &Session{
		conn: conn,
		buf:  make([]byte, bufSize),
	}
	return s
}

func (s *Session) Open() {
	log.Info().Msg("Connection opened: " + s.conn.RemoteAddr().String())
}

func (s *Session) Close() {
	s.conn.Close()
}

func (s *Session) Read() (string, error) {
	n, err := s.conn.Read(s.buf)
	if err != nil {
		if err == io.EOF {
			log.Info().Msg("Connection closed by client: " + s.conn.RemoteAddr().String())
			return "", nil
		}
		return "", err
	}

	if n > 0 {
		log.Debug().Int("readSize", n).Bytes("bytes", s.buf[:n]).Msg("Read")

		msg := bs.BytesToString(s.buf[:n])
		return msg, nil
	}

	return "", nil
}

func (s *Session) Write(message []byte) {
	log.Debug().Bytes("bytes", message).Msg("Write")

	_, err := s.conn.Write(message)
	if err != nil {
		log.Err(err).Msg("Failed to write data")
	}
}
