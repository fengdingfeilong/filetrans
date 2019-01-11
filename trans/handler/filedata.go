package handler

import (
	"encoding/binary"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/fengdingfeilong/filetrans/trans/message"
	"github.com/fengdingfeilong/filetrans/trans/message/cmdtype"

	"github.com/fengdingfeilong/roshan"
	rhandler "github.com/fengdingfeilong/roshan/handler"
	rmessage "github.com/fengdingfeilong/roshan/message"
	"github.com/fengdingfeilong/roshan/roshantool"

	"github.com/pborman/uuid"
)

type FileData struct {
	rhandler.Base
	client       *roshan.Client
	srcdir       string
	dstdir       string
	files        []*message.FileInfo
	transids     []string
	savers       map[string]*os.File
	currentIndex int
}

func NewFileData(c *roshan.Client, d string) *FileData {
	f := FileData{client: c, dstdir: d}
	f.transids = make([]string, 0)
	f.savers = make(map[string]*os.File, 0)
	return &f
}

func (h *FileData) GetBase() *rhandler.Base {
	return &h.Base
}

func (h *FileData) Execute(data []byte) {
	tid := uuid.UUID(data[0:16]).String()
	var file *message.FileInfo
	for i, v := range h.transids {
		if tid == v {
			file = h.files[i]
			break
		}
	}
	if file == nil {
		return
	}

	spath := path.Join(h.dstdir, strings.TrimLeft(file.Fullpath, h.srcdir))
	offset := binary.BigEndian.Uint64(data[16:24])
	if offset == 0 { //the first offset
		fmt.Println("first packet received")
		f, err := os.OpenFile(spath, os.O_WRONLY|os.O_CREATE, os.ModeAppend|os.ModePerm)
		if err != nil {
			f.Close()
			fmt.Println(err)
		} else {
			f.Truncate(0)
			h.savers[tid] = f
		}
	}
	// fmt.Printf("offset: %d, len: %d\n", offset, len(data)-24)
	os.MkdirAll(path.Dir(spath), os.ModePerm)
	if h.savers[tid] != nil {
		h.savers[tid].Write(data[24:])
	}
	if int64(offset)+int64(len(data))-int64(24) == file.Size { //the last offset
		fmt.Printf("last packet received, total file len is %d\n", file.Size)
		h.savers[tid].Close()
		h.sendGetFileRequest()
	}
}

func (h *FileData) Receive(para *rhandler.CommObj) {
	if para.Src == cmdtype.FileList {
		h.currentIndex = 0
		h.files = para.Obj[0].([]*message.FileInfo)
		h.srcdir = para.Obj[1].(string)
		for range h.files {
			h.transids = append(h.transids, uuid.New())
		}
		fmt.Printf("filelist received complete, total count %d\n", len(h.files))
		h.sendGetFileRequest()
	}
}

func (h *FileData) sendGetFileRequest() {
	if h.currentIndex >= len(h.files) {
		roshantool.PrintErr("handler.FileData.sendGetFileRequest", "current request index out of filelist size")
		return
	}
	msg := message.NewGetFile()
	msg.TransferID = h.transids[h.currentIndex]
	msg.Fullpath = h.files[h.currentIndex].Fullpath
	buff := rmessage.GetCommandBytes(cmdtype.GetFile, msg)
	h.Conn().Write(buff)
	h.currentIndex++
	fmt.Printf("%s transferring...\n", msg.Fullpath)
}
