package app

type commandId int

const (
	c commandId = iota
	PING
	MSG
	QUIT
)

type Command struct {
	id commandId
	client *Client
	message *Message
}
