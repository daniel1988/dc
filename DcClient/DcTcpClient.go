package main

import (
	"../TransObject"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

type DcTcpClient struct {
	addr string
}

func NewDcTcpClient(addr string) *DcTcpClient {
	return &DcTcpClient{addr}
}

func (this *DcTcpClient) Connect() *net.TCPConn {
	udpAddr, err := net.ResolveTCPAddr("tcp", this.addr)
	if err != nil {
		fmt.Println("ResolveTCPAddr Error:%v", this.addr)
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, udpAddr)
	if err != nil {
		fmt.Println("Dial Error:%v", this.addr)
		os.Exit(2)
	}
	return conn
}

func (this *DcTcpClient) send(message string) {
	conn := this.Connect()

	_, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Println("conn.Write:%v", message)
		os.Exit(3)
	}
	conn.Close()
}

func main() {
	tickerSend()
}

func tickerSend() {
	ticker := time.NewTicker(10e4)
	count := 0
	go func() {
		for range ticker.C {
			Object := TransObject.NewClient()
			Object.AppName = "HHHHHHH"

			jsonstr, _ := json.Marshal(Object)

			fmt.Println(count)
			count += 1

			udpClient := NewDcTcpClient("127.0.0.1:1024")

			udpClient.send(string(jsonstr))
		}
	}()
	time.Sleep(time.Second * 10) //阻塞，则执行次数为sleep的休眠时间/ticker的时间
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
