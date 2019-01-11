package message

import "github.com/fengdingfeilong/roshan/message"

type Disconnect struct {
	message.Message
}

func NewDisconnect() *Disconnect {
	var m Disconnect
	m.ID = ""
	m.Name = "Disconnect"
	m.Version = message.MessageVersion
	return &m
}
