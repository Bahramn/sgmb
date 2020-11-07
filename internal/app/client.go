package app

import (
	"log"
	"time"
)

type Connection interface {
	err(err error) (int, error)
	msg(msg string) (int, error)
	quit()
}

type Client struct {
	id            string
	conn          Connection
	connType      string
	lastCheckedAt time.Time
	commands      chan<- Command
}

func (c Client) ReadInput(input string) {

	msg, err := ParsMessage(input)
	if err != nil {
		c.conn.err(err)
		return
	}

	c.id = msg.Sender
	c.commands <- Command{
		id:      msg.CmdId,
		message: msg,
		client:  &c,
	}
}

func (c *Client) quit() {
	log.Printf("Quit: %s", c.id)
	c.conn.quit()
}
