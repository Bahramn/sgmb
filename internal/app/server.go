package app

import (
	"bufio"
	"fmt"
	"github.com/Bahramn/sgmb/config"
	"log"
	"net"
	"time"
)

const MegResOk = "JB_OK"
const MsgResNok = "JB_NOK"

const UnknownClient = "anonymous"

type Server struct {
	clients  map[string]*Client
	commands chan Command
}

func NewServer() *Server {
	return &Server{
		clients:  make(map[string]*Client),
		commands: make(chan Command),
	}
}

func (s *Server) Run() {
	for cmd := range s.commands {
		s.clients[cmd.client.id] = cmd.client
		cmd.client.lastCheckedAt = time.Now()

		switch cmd.id {
		case PING:
			s.ping(cmd.client, cmd.message)
		case QUIT:
			s.quit(cmd.client, cmd.message)
		case MSG:
			s.sendMsg(cmd.client, cmd.message)
		}
	}
}

func (s *Server) NewTcpClient(conn net.Conn) {
	tcpConn := &tcpConn{
		conn: conn,
	}

	c := &Client{
		conn:     tcpConn,
		id:       UnknownClient,
		commands: s.commands,
	}

	for {
		input, err := bufio.NewReader(conn).ReadString('#')
		if err != nil {
			break
		}
		c.ReadInput(input)
	}
}

func (s *Server) NewUdpClient(udp *net.UDPConn) {
	buf := make([]byte, 1024)

	for {
		n, addr, err := udp.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if addr == nil {
			fmt.Println("addr is nil")
			continue
		}

		udpConn := &udpConn{
			conn:       udp,
			udpAddress: addr,
		}

		c := &Client{
			conn:          udpConn,
			connType:      "UDP",
			id:            UnknownClient,
			lastCheckedAt: time.Now(),
			commands:      s.commands,
		}
		c.ReadInput(string(buf[0:n]))
	}
}

func (s *Server) ping(client *Client, message *Message) {
	body := message.Build()
	client.lastCheckedAt = time.Now()
	client.conn.msg(body)
}

func (s *Server) quit(client *Client, message *Message) {
	client.quit()
	delete(s.clients, client.id)
}

func (s *Server) sendMsg(client *Client, message *Message) {
	rec, ok := s.clients[message.Receiver]
	if !ok || message.Receiver == UnknownClient {
		client.conn.msg(MsgResNok)
		return
	}

	body := message.Build()
	_, err := rec.conn.msg(body)

	if err != nil {
		delete(s.clients, message.Receiver)
		client.conn.msg(MsgResNok)
	} else {
		client.conn.msg(MegResOk)
	}

}

func (s *Server) ServeTcp(conf config.Protocol) {
	if conf.Active {
		tcpListener := NewTcpServer(conf.Address)
		defer tcpListener.Close()

		for {
			conn, err := tcpListener.Accept()
			if err != nil {
				log.Printf("Unable to accept connection: %s ", err.Error())
				continue
			}
			go s.NewTcpClient(conn)
		}
	} else {
		log.Println("TCP server is not active")
	}
}

func (s *Server) ServeUdp(conf config.Protocol) {
	if conf.Active {
		udpAddr := NewUdpServer(conf.Address)
		udp, err := net.ListenUDP("udp", udpAddr)
		if err != nil {
			log.Fatalf("could not listen on UDP, " + err.Error())
		}

		go s.NewUdpClient(udp)
	} else {
		log.Println("UDP server is not active")
	}
}

func (s *Server) NumberOfClients() int {
	return len(s.clients)
}

func (s *Server) CheckClientsByLastPingAt() {
	for id, client := range s.clients {
		now := time.Now()
		if client.lastCheckedAt.Before(now.Add(-59 * time.Second)) {
			delete(s.clients, id)
		}
	}
}
