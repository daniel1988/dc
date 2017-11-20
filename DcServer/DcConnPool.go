package main

// import (
// 	Common "../DcCommon"
// 	"net"
// 	"sync"
// )

// type DcPool struct {
// 	sync.Mutex
// 	logMap map[int64]*DcConnect
// }

// func (this *DcPool) listenTcp(addr string) {
// 	defer Common.CheckPanic()

// 	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
// 	if err != nil {
// 		return fmt.Errorf("can't resolve addr : %v : %v", addr, err)
// 	}

// 	listener, err := net.ListenTCP("tcp", tcpAddr)
// 	if err != nil {
// 		return fmt.Errorf("can't listen tcp : %v : %v", addr, err)
// 	}

// 	for {
// 		conn, err := listener.AcceptTCP()
// 		if err != nil || this.quit {
// 			logErr("accept ", err, this.quit)
// 			return err
// 		}

// 		conn.SetNoDelay(true)
// 		conn.SetLinger(-1)

// 		uc := NewDcConnect(conn)
// 		go uc.Recver(this.AddConnect, this.DelConnect)
// 	}
// }

// func (this *DcPool) listenUdp(addr string) {
// 	defer Common.CheckPanic()
// }

// func (this *DcPool) listenHttp(addr string) {
// 	defer Common.CheckPanic()
// }

// // 关闭连接池
// // 停止监听并关闭所有连接
// func (this *DcPool) Close() {
// 	this.Lock()
// 	defer this.Unlock()
// 	this.quit = true
// 	for _, conn := range this.addressMap {
// 		conn.Close()
// 	}
// }

// // 将一个连接添加到连接池的一个项，但不能重复添加
// func (this *DcPool) ListAdd(list []*DcConnect, c *DcConnect) []*DcConnect {
// 	for _, v := range list {
// 		if v == c {
// 			return list
// 		}
// 	}

// 	for i, v := range list {
// 		if v == nil {
// 			list[i] = c
// 			return list
// 		}
// 	}

// 	return append(list, c)
// }

// // 将一个连接从连接池的一个项中删除
// func (this *DcPool) ListDel(list []*DcConnect, c *DcConnect) {
// 	for i, v := range list {
// 		if v == c {
// 			list[i] = nil
// 		}
// 	}
// }

// // 判断一个连接项是否为空，空的可以从池中删除
// func (this *DcPool) ListIsEmpty(list []*DcConnect) bool {
// 	for _, v := range list {
// 		if v != nil {
// 			return false
// 		}
// 	}

// 	return true
// }

// // 将连接加入连接池
// func (this *DcPool) AddConnect(uc *DcConnect) (int, int) {
// 	this.Lock()
// 	defer this.Unlock()

// 	if uc.UserID > 0 {
// 		this.userMap[uc.UserID] = this.ListAdd(this.userMap[uc.UserID], uc)
// 	}
// 	if len(uc.DeviceKey) > 0 {
// 		this.deviceMap[uc.DeviceKey] = this.ListAdd(this.deviceMap[uc.DeviceKey], uc)
// 	}
// 	this.addressMap[uc.ClientAddr] = uc

// 	return len(this.addressMap), len(this.userMap)
// }

// // 将连接从连接池删除
// func (this *DcPool) DelConnect(uc *DcConnect) (int, int) {
// 	this.Lock()
// 	defer this.Unlock()

// 	this.ListDel(this.userMap[uc.UserID], uc)
// 	if this.ListIsEmpty(this.userMap[uc.UserID]) {
// 		delete(this.userMap, uc.UserID)
// 	}
// 	this.ListDel(this.deviceMap[uc.DeviceKey], uc)
// 	if this.ListIsEmpty(this.deviceMap[uc.DeviceKey]) {
// 		delete(this.deviceMap, uc.DeviceKey)
// 	}
// 	delete(this.addressMap, uc.ClientAddr)

// 	return len(this.addressMap), len(this.userMap)
// }
