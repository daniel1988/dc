package main

import (
	Common "../DcCommon"
	DcStore "../DcStore"
	// "../TransObject"
	"encoding/json"
	"flag"
	"net"
	"os"
)

var (
	flgEs     = flag.String("eshost", "127.0.0.1", "elastic search ip")
	flgEsPort = flag.String("esport", "9200", "elastic search port")
	flgIndex  = flag.String("esindex", "dc_index", "elastic search index")
	// flgRedisSrv     = flag.String("redis", "127.0.0.1:6379", "redis server")
	flgIncrIdKey    = flag.String("es_id_key", "EsIncrIdKey", "elastic id redis key")
	flgDcCenterAddr = flag.String("dc_center", "127.0.0.1:9999", "DcCenter Address")
	flgLogDir       = flag.String("log", "/data/github/dc/Log", "log dir")
	logErr          = Common.ErrorLog
	logInfo         = Common.InfoLog
)

func init() {
	Common.Init(os.Stderr)
	Common.SetLogDir(*flgLogDir)
}

type DcCenter struct {
	*DcStore.ElasticSearch
}

func NewDcCenter() *DcCenter {
	return &DcCenter{DcStore.NewStaticStore(*flgEs, *flgEsPort)}
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

		logInfo(conn.RemoteAddr().String(), "DcSync connect success")
		this.handleConnection(conn)
	}
}

func (this *DcCenter) handleConnection(conn net.Conn) {

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		doc := make(map[string]interface{})
		json.Unmarshal(buffer[:n], &doc)

		esId := this.GenerateId()
		ttl := int(Common.NumberNow())
		err = this.InsertDoc(*flgIndex, "dc_type", esId, ttl, doc)
		if err != nil {
			logErr("InsertDoc Error:", err, string(buffer[:n]))
		}
	}
}

func (this *DcCenter) GenerateId() string {
	// RedisSrv := DcStore.NewRedisPool(*flgRedisSrv, 0)
	// esId, err := RedisSrv.Incr(*flgIncrIdKey)
	// if err != nil {
	// 	return string(Common.NumberNow())
	// }
	return Common.NewUUID()
}

func main() {
	DcCenter := NewDcCenter()
	DcCenter.listenSocket()
}
