package iface

import (
	"context"
	"net"
)

//连接接口
type Iconn interface {
	GetConnID() int    //获取链接ID
	GetConn() net.Conn //获取真实的连接对象
	Start()            //开始服务
	Stop()             //关闭连接
	//发送数据
	SendMsg(uint32, []byte)

	//获取上下文
	GetContext() context.Context

	//设置属性
	WithValue(interface{}, interface{})

	//获取配置的属性值
	Value(interface{}) interface{}
}
