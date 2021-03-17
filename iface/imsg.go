package iface

//请求数据接口
type Imsg interface {
	//获取数据id
	GetMsgId() uint32

	//获取数据长度
	GetMsgLen() uint32

	//获取数据
	GetData() []byte

	//设置数据id
	SetMsgId(id uint32)

	//设置数据长度
	SetMsgLen(msgLen uint32)

	//设置数据
	SetData(data []byte)

	String() string //打印数据
}
