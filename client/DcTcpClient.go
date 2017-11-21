package main

import (
	Common "../DcCommon"
	"encoding/json"
	"fmt"
	"net"
	"os"
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

func (this *DcTcpClient) Send(message string) {
	conn := this.Connect()

	res, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Println("conn.Write:%v", message)
		os.Exit(3)
	}
	fmt.Println(message, res)
	conn.Close()
}

func main() {
	doc := make(map[string]interface{})
	doc["DeviceKey"] = Common.NewUUID()
	doc["AppName"] = "TEST"
	doc["AppVer"] = 0
	doc["AppVerCode"] = 1.0
	doc["AppMarket"] = 1
	doc["OS"] = "xxxx"
	doc["UserID"] = int(Common.NumberNow())
	jsonstr, _ := json.Marshal(doc)

	udpClient := NewDcTcpClient("127.0.0.1:9999")

	udpClient.Send(string(jsonstr))

}
