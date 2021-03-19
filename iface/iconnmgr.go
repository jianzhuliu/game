package iface

//连接管理器
type Iconnmgr interface {
	//添加一个连接
	AddConn(conn Iconn)

	//删除一个连接
	RemoveConn(conn Iconn)

	//关闭所有连接
	ClearAllConn()

	//获取总共连接数
	GetLen() int

	//获取所有的连接
	GetAllConns() []Iconn
}
