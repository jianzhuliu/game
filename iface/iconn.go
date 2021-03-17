package iface

import "net"

//连接接口
type Iconn interface {
	GetConnID() int    //获取链接ID
	GetConn() net.Conn //获取真实的连接对象
	Start()            //开始服务
	Stop()             //关闭连接
	//发送数据
	SendMsg(uint32, []byte)
}
