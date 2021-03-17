package core

import "gitee.com/jianzhuliu/game/iface"

type Request struct {
	Connection iface.Iconn
	Message    iface.Imsg
	WorkerId   int
}

func NewRequest(conn iface.Iconn, msg iface.Imsg) iface.Irequest {
	return &Request{
		Connection: conn,
		Message:    msg,
		WorkerId:   -1,
	}
}

//获取封装连接
func (r *Request) GetConnection() iface.Iconn {
	return r.Connection
}

//获取请求数据
func (r *Request) GetData() []byte {
	return r.Message.GetData()
}

//获取数据id
func (r *Request) GetMsgId() uint32 {
	return r.Message.GetMsgId()
}

//发送数据
func (r *Request) SendMsg(msgId uint32, data []byte) {
	r.GetConnection().SendMsg(msgId, data)
}

//设置工人编号
func (r *Request) SetWorkerId(workerId int) {
	r.WorkerId = workerId
}

//获取工人编号
func (r *Request) GetWorkerId() int {
	return r.WorkerId
}
