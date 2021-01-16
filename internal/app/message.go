package app

import (
	"errors"
	"fmt"
	"github.com/Bahramn/sgmb/internal/helpers"
	"strings"
)

const MsgTypeRdy = "RDY"
const MsgTypeQuit = "QUIT"

type Message struct {
	Symbol   string
	Sender   string
	Receiver string
	Body     string
	CmdId    commandId
}

func ParsMessage(inp string) (*Message, error) {
	var err error
	var msg Message

	inp = strings.Trim(inp, "\r\n")
	args := strings.Split(inp, "*")

	if len(args) < 3 {
		err = errors.New(fmt.Sprintf("Invalid Message parts : %s", inp))
		return nil, err
	}

	msg.Sender = args[0][1:]
	msg.Receiver = args[1]
	msg.Body = helpers.TrimLastChar(args[2])
	msg.CmdId = detectCommandId(msg.Body)
	msg.Symbol = args[0][0:1]

	return &msg, err
}

func (msg *Message) Build() string {
	return fmt.Sprintf("%s%s*%s*%s#", msg.Symbol, msg.Sender, msg.Receiver, msg.Body)
}

func detectCommandId(body string) commandId {
	var id commandId

	switch body {
	case MsgTypeRdy:
		id = PING
	case MsgTypeQuit:
		id = QUIT
	default:
		id = MSG
	}

	return id
}
