package main

import (
	Common "../DcCommon"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

type UcUdpClient struct {
	addr string
}

func NewUcUdpClient(addr string) *UcUdpClient {
	return &UcUdpClient{addr}
}

func (this *UcUdpClient) Connect() *net.UDPConn {
	udpAddr, err := net.ResolveUDPAddr("udp", this.addr)
	if err != nil {
		fmt.Errorf("ResolveUDPAddr Error:%v", this.addr)
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Errorf("Dial Error:%v", this.addr)
		os.Exit(1)
	}
	return conn
}

func (this *UcUdpClient) send(message string) {
	conn := this.Connect()
	defer conn.Close()

	_, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Errorf("conn.Write:%v", message)
		os.Exit(1)
	}

	// var resp [1024]byte
	// conn.Read(resp[0:])
	// fmt.Println("Resp:", string(resp[0:]))
}

func main() {

	tickerSend()
}

func tickerSend() {
	ticker := time.NewTicker(10e6)
	count := 0
	go func() {
		for range ticker.C {

			doc := make(map[string]interface{})
			doc["DeviceKey"] = Common.NewUUID()
			doc["AppName"] = "TEST"
			doc["AppVer"] = 0
			doc["AppVerCode"] = 1.0
			doc["AppMarket"] = 1
			doc["OS"] = "xxxx"
			doc["UserID"] = int(Common.NumberNow())

			jsonstr, _ := json.Marshal(doc)
			fmt.Println(string(jsonstr))

			fmt.Println(count)
			count += 1

			udpClient := NewUcUdpClient("127.0.0.1:11110")

			udpClient.send(string(jsonstr))
		}
	}()
	time.Sleep(time.Second * 2) //阻塞，则执行次数为sleep的休眠时间/ticker的时间
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
