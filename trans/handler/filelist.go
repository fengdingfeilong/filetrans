package handler

import (
	"encoding/json"

	"github.com/fengdingfeilong/filetrans/trans/message"
	"github.com/fengdingfeilong/filetrans/trans/message/cmdtype"

	"github.com/fengdingfeilong/roshan"
	rhandler "github.com/fengdingfeilong/roshan/handler"
	"github.com/fengdingfeilong/roshan/roshantool"
)

type FileList struct {
	rhandler.Base
	client *roshan.Client
	srcdir string
	files  []*message.FileInfo
}

func NewFileList(c *roshan.Client) *FileList {
	fl := &FileList{client: c}
	fl.files = make([]*message.FileInfo, 0)
	return fl
}

func (h *FileList) GetBase() *rhandler.Base {
	return &h.Base
}

func (h *FileList) Execute(data []byte) {
	var msg message.FileList
	err := json.Unmarshal(data, &msg)
	if err != nil {
		roshantool.PrintErr("handler.FileList.Execute", err.Error())
		return
	}
	// for _, fi := range msg.Files {
	// 	fmt.Printf("md5:%s\tpath:%s\n", fi.Md5, fi.Fullpath)
	// }
	h.srcdir = msg.SrcDir
	if msg.Files == nil { //the last filelist message
		h.client.Transmit(rhandler.NewCommObj(cmdtype.FileList, cmdtype.Data, h.files, h.srcdir))
		h.files = h.files[:0]
		return
	}
	if h.srcdir != "" && h.srcdir != msg.SrcDir {
		roshantool.PrintErr("handler.FileList.Execute", "SrcDir is not the same in two groups of filelist")
	}
	h.files = append(h.files, msg.Files...)
}

func (h *FileList) Receive(para *rhandler.CommObj) {
	if para.Src == cmdtype.Accept {
		h.files = h.files[:0]
	}
}
