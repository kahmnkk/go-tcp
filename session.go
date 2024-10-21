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

func (s Session) Read() {
	for {
		n, err := s.conn.Read(s.buf)
		if err != nil {
			if err == io.EOF {
				log.Info().Msg("Connection closed by client: " + s.conn.RemoteAddr().String())
			}
			log.Err(err).Msg("Failed to receive data")
			break
		}

		if n > 0 {
			log.Debug().Int("readSize", n).Bytes("bytes", s.buf[:n]).Msg("data")
			// TODO handle packet
		}
	}
}

func (s Session) Write(message string) {
	_, err := s.conn.Write(bs.StringToBytes(message))
	if err != nil {
		log.Err(err).Msg("Failed to write data")
	}
	log.Debug().Msg("[Write] " + message)
}

func (s Session) Close() {
	s.conn.Close()
}
