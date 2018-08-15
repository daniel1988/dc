package main

import (
	Common "../DcCommon"
	"fmt"
	"net"
)

type DcTcpServer struct {
}

func NewTcpSrv() {
	return &DcTcpServer{}
}

func (this *DcTcpServer) listenSocket(addr string) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		fmt.Errorf("can't resolve addr : %v : %v", addr, err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return fmt.Errorf("can't listen tcp : %v : %v", addr, err)
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("accept ", err)
			return err
		}

		conn.SetNoDelay(true)
		conn.SetLinger(-1)

		// fmt.Println(conn.RemoteAddr().String(), " tcp connect success")

		this.handleConnection(conn)
	}
}

func (this *DcTcpServer) handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)

	// for {

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(conn.RemoteAddr().String(), " connection error: ", err)
		return
	}

	message := string(buffer[:n])

	fmt.Println("Receive:", message)
	// fmt.Println(conn.RemoteAddr().String(), "receive data string:\n", message)

	message += "\n"
	Common.FilePutContents("/tmp/tcp.log", message)
	// }
	conn.Close()

}

// func main() {
// 	DcTcpServer := NewTcpSrv()
// 	DcTcpServer.listenSocket("localhost:1024")
// }
