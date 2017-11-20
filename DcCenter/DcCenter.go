package main

import (
	Common "../DcCommon"
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	flgDcCenterAddr = flag.String("dc_center", "127.0.0.1:9999", "DcCenter Address")
	flgLogDir       = flag.String("log", "/data/dc/Log", "log dir")
	logErr          = Common.ErrorLog
	logInfo         = Common.InfoLog
)

func init() {
	Common.Init(os.Stderr)
	Common.SetLogDir(*flgLogDir)
}

type DcCenter struct {
}

func (this *DcCenter) listenSocket() error {
	addr := *flgDcCenterAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		logErr("can't resolve addr:", addr, err)
		return err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		logErr("can't listen tcp:", addr, err)
		return err
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			logErr("accept Error:", err)
			return err
		}

		conn.SetNoDelay(true)
		conn.SetLinger(-1)

		logInfo(conn.RemoteAddr().String(), " tcp connect success")

		this.handleConnection(conn)
	}
}

func (this *DcCenter) handleConnection(conn net.Conn) {

	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		fmt.Println(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
	}

}

func main() {
	DcCenter := new(DcCenter)
	DcCenter.listenSocket()
}
