package handler

import (
	"encoding/json"
	"fmt"
	"os"

	"filetrans/trans/message"

	"github.com/fengdingfeilong/roshan"
	rhandler "github.com/fengdingfeilong/roshan/handler"
	rmessage "github.com/fengdingfeilong/roshan/message"
	"github.com/fengdingfeilong/roshan/roshantool"

	"github.com/pborman/uuid"
)

const datalen int64 = 16 * 1024

type GetFile struct {
	rhandler.Base
	server *roshan.Server
}

func NewGetFile(s *roshan.Server) *GetFile {
	return &GetFile{server: s}
}

func (h *GetFile) GetBase() *rhandler.Base {
	return &h.Base
}

func (h *GetFile) Execute(data []byte) {
	var msg message.GetFile
	err := json.Unmarshal(data, &msg)
	if err != nil {
		roshantool.PrintErr("handler.GetFile.Execute", err.Error())
		return
	}
	go h.sendFile(&msg)
}

func (h *GetFile) sendFile(msg *message.GetFile) {
	fp, err := os.Open(msg.Fullpath)
	fp.Seek(msg.Offset, 0)
	defer fp.Close()
	if err != nil {
		roshantool.PrintErr("handler.GetFile.Execute", err.Error())
		return
	}
	fmt.Printf("%s transfering...\n", msg.Fullpath)
	tid := ([]byte)(uuid.Parse(msg.TransferID))
	fi, _ := fp.Stat()
	size := fi.Size()
	var buff [datalen]byte
	if size == 0 { //empty file
		d := rmessage.GetDataMsgBytes(tid, 0, buff[:0])
		h.Conn().Write(d)
		return
	}
	var c int
	for i := msg.Offset; i < size; i += int64(c) {
		c, err = fp.Read(buff[:])
		if err != nil {
			roshantool.PrintErr("handler.GetFile.Execute", err.Error())
			break
		}
		// fmt.Printf("offset: %d, len: %d\n", i, c)
		d := rmessage.GetDataMsgBytes(tid, i, buff[:c])
		_, err := h.Conn().Write(d)
		if err != nil {
			break
		}
	}
}
