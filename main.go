package main

import (
	"fmt"
	"io"
	"main/internal/bs"
	"net"
)

func handler(c net.Conn) {
	buf := make([]byte, 10)
	for {
		n, err := c.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by client: ", c.RemoteAddr().String())
			}
			fmt.Println("Failed to receive data: ", err)
			break
		}

		if n > 0 {
			fmt.Println(bs.BytesToString(buf[:n]))
			c.Write(buf[:n])
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Failed to Listen: ", err)
	}
	defer l.Close()

	fmt.Println("server running...")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Failed to Accept: ", err)
			continue
		}

		go handler(conn)
	}
}
