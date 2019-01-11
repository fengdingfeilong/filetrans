package message

import "github.com/fengdingfeilong/roshan/message"

type GetFileList struct {
	message.Message
}

func NewGetFileFile() *GetFileList {
	var m GetFileList
	m.ID = ""
	m.Name = "GetFileList"
	m.Version = message.MessageVersion
	return &m
}
