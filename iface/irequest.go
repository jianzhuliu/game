package iface

type Irequest interface {
	//获取封装连接
	GetConnection() Iconn

	//获取请求数据
	GetData() []byte

	//获取数据id
	GetMsgId() uint32

	//发送数据
	SendMsg(uint32, []byte)

	//设置工人编号
	SetWorkerId(workerId int)

	//获取工人编号
	GetWorkerId() int
}
