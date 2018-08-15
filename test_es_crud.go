package main
import (
   Common "./DcCommon"
   DcStore "./DcStore"
    "encoding/json"
    "fmt"
    "flag"
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

func main() {
    doc := make(map[string]interface{})
    doc["DeviceKey"] = Common.NewUUID()
    doc["AppName"] = "TEST"
    doc["AppVer"] = 0
    doc["AppVerCode"] = 1.0
    doc["AppMarket"] = 1
    doc["OS"] = "xxxx"
    doc["UserID"] = 1

    jsonstr, _ := json.Marshal(doc)
    fmt.Println(string(jsonstr))

     esId := Common.NewUUID()
    ttl := int(Common.NumberNow())

    es := DcStore.NewStaticStore(*flgEs, *flgEsPort)

    err := es.InsertDoc("dc_index", "dc_type", esId, ttl, jsonstr)
    if err != nil {
        panic(err)
    }
}