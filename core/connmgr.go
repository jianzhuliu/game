package core

import (
	"sync"

	"gitee.com/jianzhuliu/game/iface"
)

//连接管理器
type ConnMgr struct {
	Conns    map[int]iface.Iconn
	connLock sync.Mutex
}

//创建管理器对象
func NewConnMgr() *ConnMgr {
	return &ConnMgr{
		Conns: make(map[int]iface.Iconn),
	}
}

//添加一个连接
func (m *ConnMgr) AddConn(conn iface.Iconn) {
	m.connLock.Lock()
	defer m.connLock.Unlock()

	m.Conns[conn.GetConnID()] = conn
}

//删除一个连接
func (m *ConnMgr) RemoveConn(conn iface.Iconn) {
	m.connLock.Lock()
	defer m.connLock.Unlock()

	delete(m.Conns, conn.GetConnID())
}

//关闭所有连接
func (m *ConnMgr) ClearAllConn() {
	m.connLock.Lock()
	defer m.connLock.Unlock()

	for connId := range m.Conns {
		conn := m.Conns[connId]
		delete(m.Conns, connId)
		conn.Stop()
	}
}

//获取总共连接数
func (m *ConnMgr) GetLen() int {
	return len(m.Conns)
}

//获取所有的连接
func (m *ConnMgr) GetAllConns() []iface.Iconn {
	m.connLock.Lock()
	defer m.connLock.Unlock()

	//构建结果集
	conns := make([]iface.Iconn, 0, m.GetLen())

	//循环遍历
	for connId := range m.Conns {
		conns = append(conns, m.Conns[connId])
	}

	return conns
}
