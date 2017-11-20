package TransObject

import (
// "fmt"
// "strings"
// "time"
)

// OS类型
const (
	Android = "Android"
	IOS     = "iOS"
	PC      = "PC"
	Mobile  = "Mobile"
	HTML5   = "HTML5"
)

// 语种类型
const (
	LANG_ZH = "zh"
	LANG_EN = "en"
	LANG_HK = "hk"
)

// 应用的不可变信息
type AppSolidInfo struct {
	// 客户端设备Key
	DeviceKey string
	// 客户端appname
	AppName string
	// App版本
	AppVer int64
	// App推广渠道
	AppMarket int32
	// 操作系统
	OS string
	// 操作系统版本
	OsVer int64
	// APP 编译版本号
	AppVerCode int64
	// APP 渠道ID
	AppRefID string
}

// 客户端连接的属性,包含应用信息和用户设置
type Client struct {
	AppSolidInfo

	// 客户端的远程连接地址
	ClientAddr string
	// 客户端登录的uid
	UserID int64
	// 客户端消息开关
	Enable int64
	// 客户端语种
	Language string
	// 上线时间
	OnlineTime string
}

func NewClient() *Client {
	c := new(Client)
	c.AppName = "TEST"
	return c
}

// 传送对象
type TransObject struct {
	*Client

	// 消息ID
	MsgID string
	// 消息的来源
	SystemSrc int64

	// 生命周期，time to live
	TTL int

	// 消息事件类型
	Event string

	// 消息体
	Body []byte

	// 路由键
	RoutingKey string
}
