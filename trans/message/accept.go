package message

import (
	"filetrans/peer"
	"github.com/fengdingfeilong/roshan/message"
)

//Accept ...
type Accept struct {
	message.Message
	RemoteInfo peer.Info `json:"remoteInfo"`
}

func NewAccept() *Accept {
	var m Accept
	m.ID = ""
	m.Name = "Accept"
	m.Version = message.MessageVersion
	return &m
}
