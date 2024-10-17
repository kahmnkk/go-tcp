package main

import (
	"io"
	"main/internal/config"
	"main/internal/logger"
	"net"
	"os"

	"github.com/rs/zerolog/log"
)

func handler(c net.Conn) {
	buf := make([]byte, 10)
	for {
		n, err := c.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Info().Msg("Connection closed by client: " + c.RemoteAddr().String())
			}
			log.Err(err).Msg("Failed to receive data")
			break
		}

		if n > 0 {
			log.Debug().Int("readSize", n).Bytes("bytes", buf[:n]).Msg("data")
			c.Write(buf[:n])
		}
	}
}

func main() {
	// Config
	conf := config.Read()

	// Logger
	log.Logger = logger.NewLogger(conf.LogLevel, os.Stdout)

	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Err(err).Msg("Failed to Listen")
	}
	defer l.Close()

	log.Info().Msg("Sever running...")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Err(err).Msg("Failed to Accept")
			continue
		}

		go handler(conn)
	}
}
