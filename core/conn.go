package core

import (
	"context"
	"net"

	"gitee.com/jianzhuliu/game/iface"
)

//连接实现对象
type Connection struct {
	Server     iface.Iserver      //服务对象
	Conn       net.Conn           //真实的连接
	ConnId     int                //连接id
	Ctx        context.Context    //上下文
	CancelFunc context.CancelFunc //上下文取消函数
	isClosed   bool               //服务器是否关闭标识

	ChanMsg chan iface.Imsg //请求响应,读写分离模式
}

//构建连接对象
func NewConnection(conn net.Conn, connId int, s iface.Iserver) *Connection {
	connection := &Connection{
		Server:  s,
		Conn:    conn,
		ConnId:  connId,
		ChanMsg: make(chan iface.Imsg),
	}

	//添加到连接管理器中
	s.GetConnMgr().AddConn(connection)

	return connection
}

//开启读操作
func (c *Connection) startReader() {
	Logger.Printf("[Connection.startReader] [%s] connId=%d ==========>", c.Conn.RemoteAddr(), c.GetConnID())
	defer c.Stop()
	for {
		select {
		case <-c.Ctx.Done():
			return
		default:
			//解包数据
			dbpack := NewDbpack()
			msg, err := dbpack.Unpack(c.Conn)
			if err != nil {
				return
			}

			//构造请求
			request := NewRequest(c, msg)

			//如果开启了工作池模式，需要发送到工作池
			if ConfObj.WorkerPoolSize > 0 {
				c.Server.GetRouter().SendRequestToWorkerPool(request)
			} else {
				go c.Server.GetRouter().DoHandle(request)
			}
		}
	}
}

//开启写操作
func (c *Connection) startWriter() {
	Logger.Printf("[Connection.startWriter] [%s] connId=%d ==========>", c.Conn.RemoteAddr(), c.GetConnID())

	for {
		select {
		case <-c.Ctx.Done():
			return
		case msg := <-c.ChanMsg:
			dbpack := NewDbpack()
			binaryMsg, err := dbpack.Pack(msg)
			if err != nil {
				Logger.Printf("[Connection.startWriter] pack msg(%v) fail, err=%v", msg, err)
				continue
			}

			if _, err := c.GetConn().Write(binaryMsg); err != nil {
				Logger.Printf("[Connection.startWriter] send msg(%v) fail, err=%v", msg, err)
				continue
			}
		}
	}
}

//发送数据
func (c *Connection) SendMsg(msgId uint32, data []byte) {
	//发送消息到管道
	msg := NewMessage(msgId, data)
	c.ChanMsg <- msg
}

//连接开始服务
func (c *Connection) Start() {
	Logger.Print("[Connection.Start] ==========>")
	c.Ctx, c.CancelFunc = context.WithCancel(context.Background())
	go c.startReader()
	go c.startWriter()
}

//关闭服务器
func (c *Connection) Stop() {
	if c.isClosed {
		return
	}

	Logger.Print("[Connection.Stop] ==========>", c.GetConnID())

	//关闭新连接
	c.CancelFunc()
	c.isClosed = true
	close(c.ChanMsg)

	c.GetConn().Close()
}

//获取链接ID
func (c *Connection) GetConnID() int {
	return c.ConnId
}

//获取真实的连接对象
func (c *Connection) GetConn() net.Conn {
	return c.Conn
}