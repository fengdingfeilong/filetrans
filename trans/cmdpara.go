package trans

//CmdPara ...
type CmdPara struct {
	Help    bool
	LogName string
	//Source source directory
	Source string
	//Target target directory
	Target string
	//Key incrypt the temp pubkey
	Key string
	//IP source ip
	IP string
}
