package main

import (
	Common "./DcCommon"
	"./DcStore"
	// "./TransObject"
	"fmt"
	"log"
)

var ()

func main() {
	RedisSrv := DcStore.NewRedisPool("127.0.0.1:6379", 0)

	RedisSrv.SetEx("foo", "hello", 86400)
	fmt.Println(RedisSrv.Get("foo"))

	res, err := RedisSrv.Rpush("lfoo", []byte(Common.FormatNow()))
	fmt.Println(res, err)
	message, _ := RedisSrv.Lpop("lfoo")

	fmt.Println(string(message))

	fmt.Println(Common.NumberNow())
}

func Log(v ...interface{}) {
	log.Println(v...)
}
