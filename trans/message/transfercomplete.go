package message

import (
	"github.com/fengdingfeilong/roshan/message"
)

//TransferComplete ...
type TransferComplete struct {
	message.Message
}

func NewTransferComplete() *TransferComplete {
	var m TransferComplete
	m.ID = ""
	m.Name = "TransferComplete"
	m.Version = message.MessageVersion
	return &m
}
