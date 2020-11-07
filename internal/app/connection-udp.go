package app

import (
	"log"
	"net"
)

type udpConn struct {
	conn       *net.UDPConn
	udpAddress *net.UDPAddr
}

func NewUdpServer(addr string) *net.UDPAddr {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatalf("could not resolve UDP addr, " + err.Error())
	}

	log.Printf("UDP server started on %s", addr)

	return udpAddr
}

func (u *udpConn) msg(msg string) (int, error) {
	log.Printf("Message udp: %s", msg)
	return u.conn.WriteTo([]byte(msg+"\n"), u.udpAddress)
}

func (u *udpConn) err(err error) (int, error) {
	log.Printf("Message: %s", err.Error())
	return u.conn.WriteTo([]byte(err.Error()+"\n"), u.udpAddress)
}

func (u *udpConn) quit() {
	u.conn.Close()
}
