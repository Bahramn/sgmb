package app

import (
	"log"
	"net"
)

type tcpConn struct {
	conn net.Conn
}

func NewTcpServer(addr string) net.Listener {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Unable to start server: %s", err.Error())
	}

	log.Printf("TCP server started on %s", addr)

	return listener
}

func (t *tcpConn) msg(msg string) (int, error) {
	log.Printf("Message tcp: %s", msg)
	return t.conn.Write([]byte(msg + "\n"))
}
func (t *tcpConn) err(err error) (int, error) {
	log.Printf("Message: %s", err.Error())
	return t.conn.Write([]byte(err.Error() + "\n"))
}

func (t *tcpConn) quit() {
	t.conn.Close()
}
