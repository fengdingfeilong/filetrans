package handler

import (
	"encoding/json"
	"fmt"

	"github.com/fengdingfeilong/filetrans/trans/message"
	"github.com/fengdingfeilong/filetrans/trans/message/cmdtype"

	"github.com/fengdingfeilong/roshan"
	"github.com/fengdingfeilong/roshan/roshantool"

	rhandler "github.com/fengdingfeilong/roshan/handler"
	rmessage "github.com/fengdingfeilong/roshan/message"
)

type Accept struct {
	rhandler.Base
	client *roshan.Client
}

func NewAccept(c *roshan.Client) *Accept {
	return &Accept{client: c}
}

func (h *Accept) GetBase() *rhandler.Base {
	return &h.Base
}

func (h *Accept) Execute(data []byte) {
	var msg message.Accept
	err := json.Unmarshal(data, &msg)
	if err != nil {
		roshantool.PrintErr("handler.Accept.Execute", err.Error())
	} else {
		roshantool.Printf("%s accept", msg.RemoteInfo.Addr)
		fmt.Println("connected, begin transfer")
		h.client.Transmit(rhandler.NewCommObj(cmdtype.Accept, cmdtype.FileList))
		m := message.NewGetFile()
		buff := rmessage.GetCommandBytes(cmdtype.GetFileList, m)
		h.Conn().Write(buff)
	}
}
