package message

import (
	"github.com/fengdingfeilong/roshan/message"
)

type FileInfo struct {
	Name     string `json:"name"`
	Fullpath string `json:"fullpath"`
	Size     int64  `json:"size"`
	Md5      string `json:"MD5"`
}

type FileList struct {
	message.Message
	SrcDir string      `json:"srcdir"`
	Files  []*FileInfo `json:"files"`
}

func NewFileList() *FileList {
	var m FileList
	m.ID = ""
	m.Name = "FileList"
	m.Version = message.MessageVersion
	return &m
}
