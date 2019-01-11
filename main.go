package main

import (
	"bufio"
	"filetrans/trans"
	"filetrans/trans/handler"
	"filetrans/trans/message"
	"filetrans/trans/message/cmdtype"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/fengdingfeilong/roshan"
	rmessage "github.com/fengdingfeilong/roshan/message"
	"github.com/fengdingfeilong/roshan/roshantool"
)

var para *trans.CmdPara

func main() {
	para = getCmdPara()
	if showHelp(para) {
		return
	}
	roshantool.CreateLog(para.LogName)
	roshantool.InnerLog = roshanlog

	isser, es := isServer(para)
	iscli := false
	if !isser {
		iscli, es = isClient(para)
	}

	if !isser && !iscli {
		fmt.Printf("argument missing\n%s", es)
		showHelp(para)
		return
	}

	if isser {
		startServer(para)
	} else if iscli {
		startClient(para)
	}

	bufio.NewReader(os.Stdin).ReadByte()
	if isser {
		//close all client here
		server.CloseListen()
	}
	if iscli {
		client.Close()
	}

	roshantool.CloseLog()
}

func getCmdPara() *trans.CmdPara {
	var para trans.CmdPara
	flag.BoolVar(&para.Help, "h", false, "show help")
	flag.StringVar(&para.LogName, "Log", "filetrans.log", "Log File")
	flag.StringVar(&para.Source, "Source", "", "Folder To Transfer")
	flag.StringVar(&para.Target, "Target", "", "Folder to place data. Folder must exist and be empty.if not, error out")
	flag.StringVar(&para.Key, "Key", "", "Key")
	flag.StringVar(&para.IP, "IP", "", "source ip")
	flag.Parse()
	return &para
}

func isServer(para *trans.CmdPara) (bool, string) {
	_, err := os.Stat(para.Source)
	if err != nil {
		var s string
		if os.IsNotExist(err) {
			s = "source is not exist"
		} else {
			s = err.Error()
		}
		return false, s
	}
	return true, ""
}

func isClient(para *trans.CmdPara) (bool, string) {
	_, err := os.Stat(para.Target)
	if err != nil {
		var s string
		if os.IsNotExist(err) {
			s = "target is not exist"
		} else {
			s = err.Error()
		}
		return false, s
	}
	if para.IP == "" {
		return false, "source IP can not be empty"
	}
	if para.Key == "" {
		return false, "Key can not be empty"
	}
	return true, ""
}

var port = 8102

var server *roshan.Server

func startServer(para *trans.CmdPara) {
	server = roshan.NewServer()
	server.BeforeAccept = beforeAccept
	server.SocketAccepted = servAccepted
	server.CmdMessageReceived = cmdMsgReceived
	server.AddHandler(cmdtype.Connect, handler.NewConnect(server))
	server.AddHandler(cmdtype.Ping, handler.NewPing())
	server.AddHandler(cmdtype.Disconnect, handler.NewDisconnect(server))
	server.AddHandler(cmdtype.GetFileList, handler.NewGetFileList(server, para.Source))
	server.AddHandler(cmdtype.GetFile, handler.NewGetFile(server))
	server.Start(port)
}

var client *roshan.Client

func startClient(para *trans.CmdPara) {
	client = roshan.NewClient()
	client.SocketConnected = cliConnected
	client.BeforeClose = cliBeforeClose
	client.CmdMessageReceived = cmdMsgReceived
	client.AddHandler(cmdtype.Accept, handler.NewAccept(client))
	client.AddHandler(cmdtype.Reject, handler.NewReject(client))
	client.AddHandler(cmdtype.FileList, handler.NewFileList(client))
	client.AddHandler(cmdtype.Data, handler.NewFileData(client, para.Target))
	client.Connect(para.IP, port)
}

func showHelp(para *trans.CmdPara) bool {
	if para.Help {
		flag.PrintDefaults()
		return true
	}
	return false
}

func beforeAccept() {

}

func servAccepted(conn net.Conn) {
	roshantool.Println("accept socket :" + conn.RemoteAddr().String())
	server.StopHandlePacket(conn) //stop handle command and data packet until accept connectmessage
}

func cliConnected(conn net.Conn) {
	roshantool.Println("connected socket :" + conn.RemoteAddr().String())

	m := message.NewConnect()
	m.RemoteInfo.Addr = conn.LocalAddr().String()

	buff := rmessage.GetCommandBytes(cmdtype.Connect, m)
	conn.Write(buff)
}

func cliBeforeClose(conn net.Conn) {
	m := message.NewDisconnect()
	buff := rmessage.GetCommandBytes(cmdtype.Disconnect, m)
	conn.Write(buff)
	roshantool.Println("client close the connection")
}

func cmdMsgReceived(conn net.Conn, t rmessage.CmdType) {
	roshantool.Println("received cmd message :" + cmdtype.GetCmdString(t))
	if t == cmdtype.Connect {
		server.StartHandlePacket(conn)
	}
}

func roshanlog(info string, err error) {
	roshantool.Println(info)
	if err != nil {
		roshantool.Println(err.Error())
	}
}
