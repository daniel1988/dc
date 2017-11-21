package main

import (
	Common "../DcCommon"
	"flag"
	"os"
)

var (
	flgDcCenterAddr  = flag.String("dc_center", "127.0.0.1:9999", "DcCenter Address")
	flgRedisSrv      = flag.String("redis", "127.0.0.1:6379", "redis server")
	flgUdpAddr       = flag.String("udp_srv", "127.0.0.1:11110", "udp server address")
	flgRedisQueueKey = flag.String("redisqueue", "RedisQueueKey", "redis queue key")
	flgLogDir        = flag.String("log", "/data/dc/Log", "log dir")
	logErr           = Common.ErrorLog
	logInfo          = Common.InfoLog
)

func init() {
	Common.Init(os.Stderr)
	Common.SetLogDir(*flgLogDir)
}

func ListenUdpSrv() {
	udpSrv := new(DcUdpServer)
	udpSrv.ListenSocket()
}

func StartDcSync() {
	syncSrv := NewDcSync()
	for {
		syncSrv.Sync()
	}
}

func main() {
	defer Common.CheckPanic()

	go StartDcSync()

	ListenUdpSrv()
}
