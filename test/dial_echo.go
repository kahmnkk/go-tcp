package test

import (
	"io"
	"main/internal/bs"
	"net"
	"testing"
)

func TestEcho(t *testing.T) {
	// 소켓 bind + listen
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		defer func() {
			done <- struct{}{}
		}()

		for {
			// 연결 수립
			conn, err := listener.Accept()
			if err != nil {
				t.Log(err)
				return
			}

			// Handler. 고루틴으로 각 연결 처리
			go func(c net.Conn) {
				defer func() {
					c.Close()
					done <- struct{}{}
				}()

				buf := make([]byte, 10)
				for {
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}

					t.Logf("received: %q", buf[:n])
					c.Write(buf[:n])
				}
			}(conn)
		}
	}()

	t.Log(listener.Addr().String())

	// 클라이언트측 연결
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Write([]byte(bs.StringToBytes("Hello World!")))
	if err != nil {
		if nErr, ok := err.(net.Error); ok && nErr.Temporary() {
			t.Logf("temporary error: %v", nErr)
		}
		t.Error(err)
	}

	// 연결 종료
	conn.Close()
	<-done

	listener.Close()
	<-done
}
