package peer

import (
	"github.com/fengdingfeilong/filetrans/os"
)

type Info struct {
	Addr string `json:"addr"`
	OS   os.OS  `json:"os"`
}
