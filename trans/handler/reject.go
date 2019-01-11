package handler

import (
	"encoding/json"
	"filetrans/trans/message"

	"github.com/fengdingfeilong/roshan"
	rhandler "github.com/fengdingfeilong/roshan/handler"
	"github.com/fengdingfeilong/roshan/roshantool"
)

type Reject struct {
	rhandler.Base
	client *roshan.Client
}

func NewReject(c *roshan.Client) *Reject {
	return &Reject{client: c}
}

func (h *Reject) GetBase() *rhandler.Base {
	return &h.Base
}

func (h *Reject) Execute(data []byte) {
	var msg message.Reject
	err := json.Unmarshal(data, &msg)
	if err != nil {
		roshantool.PrintErr("handler.Reject.Execute", err.Error())
	} else {
		roshantool.Print("%s Reject, reason: %s ", msg.RemoteInfo.Addr, msg.Reason)
	}
}
