package main

import (
	Common "../DcCommon"
	"../DcStore"
	// "flag"
	"fmt"
	"net"
	"os"
	"time"
)

var (
	// flgDcCenterAddr = flag.String("dc_center", ":9999", "DcCenter Address")
	// flgRedisSrv      = flag.String("redis", "127.0.0.1:6379", "redis server")
	// flgRedisQueueKey = flag.String("redisqueue", "RedisQueueKey", "redis queue key")

	// flgLogDir    = flag.String("log", "/data/dc/Log", "log dir")
	// logErr       = Common.ErrorLog
	// logInfo      = Common.InfoLog
	messageCount = 0
)

func init() {
	Common.Init(os.Stderr)
	Common.SetLogDir(*flgLogDir)
}

type DcSync struct {
	addr string
}

func NewDcSync() *DcSync {
	return &DcSync{*flgDcCenterAddr}
}

func (this *DcSync) Connect() *net.TCPConn {
	udpAddr, err := net.ResolveTCPAddr("tcp", this.addr)
	if err != nil {
		logErr("ResolveTCPAddr Error:", this.addr)
		return nil
	}
	conn, err := net.DialTCP("tcp", nil, udpAddr)
	if err != nil {
		logErr("Dial Error:", this.addr)
		return conn
	}
	return conn
}

func (this *DcSync) Close(conn *net.TCPConn) error {
	conn.Close()
	return nil
}

func (this *DcSync) Sync() {
	conn := this.Connect()

	RedisSrv := DcStore.NewRedisPool(*flgRedisSrv, 0)
	message, err := RedisSrv.Lpop(*flgRedisQueueKey)
	if message == nil {
		fmt.Print(".", messageCount)
		time.Sleep(time.Second * 10)
		messageCount = 0
		return
	}
	messageCount += 1
	if err != nil {
		logErr("Lpop Error", *flgRedisQueueKey)
		return
	}

	_, err = conn.Write(message)
	if err != nil {
		logErr("Re Rpush", *flgRedisQueueKey, string(message))
		RedisSrv.Rpush(*flgRedisQueueKey, message)
		return
	}
	fmt.Println("Sync:", string(message))
	conn.Close()
}

// func main() {
// 	syncSrv := NewDcSync(*flgDcCenterAddr)
// 	for {
// 		syncSrv.Sync()
// 	}

// }
