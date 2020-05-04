package handler

import (
	"strings"

	"filetrans/trans/message/cmdtype"

	rhandler "github.com/fengdingfeilong/roshan/handler"
	"github.com/fengdingfeilong/roshan/message"
	"github.com/fengdingfeilong/roshan/roshantool"
)

type Ping struct {
	rhandler.Base
}

func NewPing() *Ping {
	return &Ping{}
}

func (h *Ping) GetBase() *rhandler.Base {
	return &h.Base
}

func (h *Ping) Execute(data []byte) {
	if strings.ToUpper(string(data)) == "PING" {
		buff := message.GetCommandBytes(cmdtype.Ping, "PONG")
		h.Conn().Write(buff)
	} else if strings.ToUpper(string(data)) == "PONG" {
		roshantool.Println("received PONG")
	} else {
		roshantool.PrintErr("handler.ping.Execute", "received unnormal data: "+string(data))
	}
}
