package handler

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/fengdingfeilong/filetrans/trans/message"
	"github.com/fengdingfeilong/filetrans/trans/message/cmdtype"

	"github.com/fengdingfeilong/roshan"
	rhandler "github.com/fengdingfeilong/roshan/handler"
	rmessage "github.com/fengdingfeilong/roshan/message"
	"github.com/fengdingfeilong/roshan/roshantool"
)

type GetFileList struct {
	rhandler.Base
	server *roshan.Server
	dir    string
}

func NewGetFileList(s *roshan.Server, d string) *GetFileList {
	return &GetFileList{server: s, dir: d}
}

func (h *GetFileList) GetBase() *rhandler.Base {
	return &h.Base
}

func (h *GetFileList) Execute(data []byte) {
	var msg message.GetFileList
	err := json.Unmarshal(data, &msg)
	if err != nil {
		roshantool.PrintErr("handler.GetFileList.Execute", err.Error())
		return
	}
	files := getFileList(h.dir)
	for i := 0; i < len(files); i += 50 {
		j := i + 50
		if j > len(files) {
			j = len(files)
		}
		m := message.NewFileList()
		m.SrcDir = h.dir
		m.Files = files[i:j]
		buff := rmessage.GetCommandBytes(cmdtype.FileList, m)
		h.Conn().Write(buff)
	}
	m := message.NewFileList()
	m.SrcDir = h.dir
	buff := rmessage.GetCommandBytes(cmdtype.FileList, m)
	h.Conn().Write(buff)
}

func getFileList(dir string) []*message.FileInfo {
	files := make([]*message.FileInfo, 0)
	rd, err := ioutil.ReadDir(dir)
	if err != nil {
		roshantool.PrintErr("handler.GetFileList.getFileList", err.Error())
		return nil
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fs := getFileList(path.Join(dir, fi.Name()))
			if fs != nil {
				files = append(files, fs...)
			}
		} else {
			var file message.FileInfo
			file.Name = fi.Name()
			file.Size = fi.Size()
			file.Fullpath = path.Join(dir, file.Name)
			file.Md5 = getFileMD5(file.Fullpath)
			files = append(files, &file)
		}
	}
	return files
}

func getFileMD5(path string) string {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		roshantool.PrintErr("handler.GetFileList.getFileMD5", err.Error())
		return ""
	}
	md5 := md5.New()
	_, err = io.Copy(md5, file)
	if err != nil {
		roshantool.PrintErr("handler.GetFileList.getFileMD5", err.Error())
		return ""
	}
	s := hex.EncodeToString(md5.Sum(nil))
	return s
}
