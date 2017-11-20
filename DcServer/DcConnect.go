package main

// import (
// 	"../TransObject"
// 	"net"
// 	"sync"
// 	"time"
// )

// type DcConnect struct {
// 	net.Conn
// 	*TransObject.Client
// 	msgRecved []string
// 	lastReqID int
// 	lastReqTm time.Time
// 	lastReq   string
// }

// func NewDcConnect(socket net.Conn) *DcConnect {
// 	client := TransObject.NewClient(socket.RemoteAddr().String())
// 	return &DcConnect{socket, client, nil, 0, time.Now(), ""}
// }
