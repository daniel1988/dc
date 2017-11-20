package main

import (
	Common "../DcCommon"
	"../DcStore"
	"fmt"
	"net"
	"os"
)

var (
// flgRedisSrv      = flag.String("redis", "127.0.0.1:6379", "redis server")
// flgUdpAddr       = flag.String("udp_srv", "127.0.0.1:11110", "udp server address")
// flgRedisQueueKey = flag.String("redisqueue", "RedisQueueKey", "redis queue key")

// flgLogDir        = flag.String("log", "/data/dc/Log", "log dir")
// logErr           = Common.ErrorLog
// logInfo          = Common.InfoLog
)

func init() {
	Common.Init(os.Stderr)
	Common.SetLogDir(*flgLogDir)
}

type DcUdpServer struct {
}

func NewUdpSrv() *DcUdpServer {
	return &DcUdpServer{}
}

func (this *DcUdpServer) ListenSocket() {
	udpAddr, err := net.ResolveUDPAddr("udp", *flgUdpAddr)
	if err != nil {
		logErr("ResolveUDPAddr:", err)
		return
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		logErr("ListenUDP:", err)
		return
	}
	fmt.Println("UDPSrv Started At Addr:", *flgUdpAddr)
	for {
		this.HandleConnection(conn)
	}

}

func (this *DcUdpServer) HandleConnection(conn *net.UDPConn) {
	var buf [1024]byte

	n, raddr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		logErr("ReadFromUDP:", err)
		return
	}

	message := string(buf[0:n])
	logInfo("Receive: ", message)

	RedisSrv := DcStore.NewRedisPool(*flgRedisSrv, 0)
	_, err = RedisSrv.Rpush(*flgRedisQueueKey, []byte(message))
	if err != nil {
		logErr("Rpush Error: ", *flgRedisQueueKey, message)
		return
	}

	_, err = conn.WriteToUDP([]byte("msg Received!"), raddr)
	if err != nil {
		logErr("UDP Resp Error", err)
	}
}

// func main() {

// 	udpSrv := new(DcUdpServer)
// 	udpSrv.ListenSocket()

// }
