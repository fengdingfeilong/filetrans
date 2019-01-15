package message

import (
	"github.com/fengdingfeilong/roshan/message"
)

type GetFile struct {
	message.Message
	TransferID string `json:"transferID"`
	Fullpath   string `json:"fullpath"`
	Offset     int64  `Offset`
}

func NewGetFile() *GetFile {
	var m GetFile
	m.ID = ""
	m.Name = "GetFile"
	m.Version = message.MessageVersion
	return &m
}
