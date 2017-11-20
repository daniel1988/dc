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
	conn *net.TCPConn
}

func NewDcSync() *DcSync {
	return &DcSync{*flgDcCenterAddr, nil}
}

func (this *DcSync) Connect() *net.TCPConn {
	if this.conn != nil {
		return this.conn
	}
	udpAddr, err := net.ResolveTCPAddr("tcp", this.addr)
	if err != nil {
		this.conn = nil
		logErr("ResolveTCPAddr Error:", this.addr)
		return this.conn
	}
	this.conn, err = net.DialTCP("tcp", nil, udpAddr)
	if err != nil {
		this.conn = nil
		logErr("Dial Error:", this.addr)
		return this.conn
	}
	return this.conn
}

func (this *DcSync) Close() error {
	this.conn.Close()
	this.conn = nil
	return nil
}

func (this *DcSync) Sync() {
	this.Connect()
	conn := this.conn

	RedisSrv := DcStore.NewRedisPool(*flgRedisSrv, 0)
	message, err := RedisSrv.Lpop(*flgRedisQueueKey)
	if message == nil {
		fmt.Print(".", messageCount)
		time.Sleep(time.Second * 10)
		messageCount = 0
		return
	}
	messageCount += 1
	fmt.Println("Lpop:", string(message))
	if err != nil {
		logErr("Lpop Error", *flgRedisQueueKey)
		return
	}

	_, err = conn.Write(message)
	if err != nil {
		logErr("Re Rpush", *flgRedisQueueKey, message)
		RedisSrv.Rpush(*flgRedisQueueKey, message)
		return
	}
	fmt.Println("Sync:", string(message))
	// conn.Close()
}

// func main() {
// 	syncSrv := NewDcSync(*flgDcCenterAddr)
// 	for {
// 		syncSrv.Sync()
// 	}

// }
