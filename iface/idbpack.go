package iface

import "io"

//数据解包装包接口
type Ipack interface {
	//获取头部数据长度
	GetHeaderLen() int

	//打包
	Pack(Imsg) ([]byte, error)

	//拆包
	Unpack(io.Reader) (Imsg, error)
}
