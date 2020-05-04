package handler

import (
	"encoding/json"

	"filetrans/trans/message"
	"filetrans/trans/message/cmdtype"

	"github.com/fengdingfeilong/roshan"
	rhandler "github.com/fengdingfeilong/roshan/handler"
	rmessage "github.com/fengdingfeilong/roshan/message"
	"github.com/fengdingfeilong/roshan/roshantool"
)

//Connect ...
type Connect struct {
	rhandler.Base
	server *roshan.Server
}

func NewConnect(s *roshan.Server) *Connect {
	return &Connect{server: s}
}

func (h *Connect) GetBase() *rhandler.Base {
	return &h.Base
}

func (h *Connect) Execute(data []byte) {
	var msg message.Connect
	err := json.Unmarshal(data, &msg)
	if err != nil {
		roshantool.PrintErr("handler.Connect.Execute", err.Error())
	}
	//if reject  h.Server.StopHandlePacket(h.Conn())
	m := message.NewAccept()
	m.RemoteInfo.Addr = h.Conn().LocalAddr().String()

	buff := rmessage.GetCommandBytes(cmdtype.Accept, m)
	h.Conn().Write(buff)
	h.server.StartHandlePacket(h.Conn())
}
