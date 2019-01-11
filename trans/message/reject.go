package message

import (
	"github.com/fengdingfeilong/filetrans/peer"
	"github.com/fengdingfeilong/roshan/message"
)

//Reject ...
type Reject struct {
	message.Message
	RemoteInfo peer.Info `json:"remoteInfo"`
	Reason     string    `json:"reason"`
}

func NewReject() *Reject {
	var m Reject
	m.ID = ""
	m.Name = "Reject"
	m.Version = message.MessageVersion
	return &m
}
