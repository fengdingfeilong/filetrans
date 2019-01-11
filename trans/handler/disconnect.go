package handler

import (
	"encoding/json"

	"github.com/fengdingfeilong/filetrans/trans/message"

	"github.com/fengdingfeilong/roshan"
	rhandler "github.com/fengdingfeilong/roshan/handler"
	"github.com/fengdingfeilong/roshan/roshantool"
)

type Disconnect struct {
	rhandler.Base
	server *roshan.Server
}

func NewDisconnect(s *roshan.Server) *Disconnect {
	return &Disconnect{server: s}
}

func (h *Disconnect) GetBase() *rhandler.Base {
	return &h.Base
}

func (h *Disconnect) Execute(data []byte) {
	var msg message.Disconnect
	err := json.Unmarshal(data, &msg)
	if err != nil {
		roshantool.PrintErr("handler.Disconnect.Execute", err.Error())
	}
	//do something such as notify
}
