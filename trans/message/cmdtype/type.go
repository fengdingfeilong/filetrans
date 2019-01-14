package cmdtype

import (
	"strconv"

	"github.com/fengdingfeilong/roshan/message"
)

const (
	//Data this is a special cmdtype for data message, the value must be 0
	Data    = 0
	Connect = message.CmdType(iota + 0x2B)
	Accept
	Reject
	Disconnect
	Ping
	Transfer
	TransferCancel
	TransferComplete
	Error
	Command

	GetFileList = message.CmdType(iota + 0x012B)
	FileList
	GetFile
)

var typedes = [...]string{
	"Connect",
	"Accept",
	"Reject",
	"Disconnect",
	"Ping",
	"Transfer",
	"TransferCancel",
	"TransferSuccess",
	"Error",
	"Command",
}

var typedes2 = [...]string{
	"GetFileList",
	"FileList",
	"GetFile",
}

func GetCmdString(t message.CmdType) string {
	if t >= Connect && t < GetFileList {
		i := int(t - Connect)
		if i < len(typedes) {
			return typedes[i]
		}
		return strconv.Itoa(int(t))
	}
	if t >= GetFileList {
		i := int(t - GetFileList)
		if i < len(typedes2) {
			return typedes2[i]
		}
		return strconv.Itoa(int(t))
	}
	return strconv.Itoa(int(t))
}
