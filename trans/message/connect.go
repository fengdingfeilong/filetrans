package message

import (
	"github.com/fengdingfeilong/filetrans/peer"
	"github.com/fengdingfeilong/roshan/message"
)

//Connect ...
type Connect struct {
	message.Message
	RemoteInfo peer.Info `json:"remoteInfo"`
}

func NewConnect() *Connect {
	var m Connect
	m.ID = ""
	m.Name = "Connect"
	m.Version = message.MessageVersion
	return &m
}
