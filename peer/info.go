package peer

import (
	"filetrans/os"
)

type Info struct {
	Addr string `json:"addr"`
	OS   os.OS  `json:"os"`
}
