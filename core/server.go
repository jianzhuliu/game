package core

import (
	"fmt"
	"net"

	"gitee.com/jianzhuliu/game/iface"
)

//服务器具体实现
type Server struct {
	Network string         //服务器协议类型，tcp tcp4 tcp6
	Host    string         //服务器ip
	Port    int            //监听端口号
	Router  iface.Irouter  //消息路由
	ConnMgr iface.Iconnmgr //连接管理器

	OnConnStartFunc func(iface.Iconn) //建立链接时的钩子方法
	OnConnStopFunc  func(iface.Iconn) //链接退出时的钩子方法
}

func NewServer(host string, port int) iface.Iserver {
	return &Server{
		Network: ConfObj.Network,
		Host:    host,
		Port:    port,
		Router:  NewHandlers(),
		ConnMgr: NewConnMgr(),
	}
}

//获取管理器对象
func (s *Server) GetConnMgr() iface.Iconnmgr {
	return s.ConnMgr
}

//获取路由处理对象
func (s *Server) GetRouter() iface.Irouter {
	return s.Router
}

//添加消息id对应处理对象
func (s *Server) AddRouter(msgId uint32, handler iface.Ihandler) {
	s.Router.AddRouter(msgId, handler)
}

//开始服务
func (s *Server) start() {
	Logger.Print("[Server.start] ==========>")
	//监听端口
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	ln, err := net.Listen(s.Network, address)
	if err != nil {
		Logger.Fatalf("listen address(%s) error=%v", address, err)
	}

	defer ln.Close()
	defer s.Stop()

	//开启工作池
	s.Router.OpenWorkerPool()

	var id int

	for {
		//循环监听连接
		conn, err := ln.Accept()
		if err != nil {
			Logger.Print("accept err=", err)
			continue
		}

		//连接数判断
		if s.ConnMgr.GetLen() > ConfObj.MaxConn {
			Logger.Printf("[%s]连接数超出上限(%d)", conn.RemoteAddr(), ConfObj.MaxConn)
			conn.Close()
			continue
		}

		id++
		//封装连接
		connection := NewConnection(conn, id, s)

		go connection.Start()
	}
}

//启动入口
func (s *Server) Run() {
	Logger.Print("[Server.run] ==========>")
	Logger.Print("[Server.run|conf]", ConfObj)

	go s.start()

	//阻塞
	for {
	}
}

//关闭服务器
func (s *Server) Stop() {
	Logger.Print("[Server.Stop] ==========>")

	//关闭所有连接
	s.GetConnMgr().ClearAllConn()
}

//设置日志对象
func (s *Server) SetLogger(logger iface.Ilogger) {
	Logger.Print("[Server.SetLogger] ==========>")
	Logger = logger
}

//设置连接建立时钩子方法
func (s *Server) SetOnConnStartFunc(fn func(iface.Iconn)) {
	s.OnConnStartFunc = fn
}

//设置连接关闭时钩子方法
func (s *Server) SetOnConnStopFunc(fn func(iface.Iconn)) {
	s.OnConnStopFunc = fn
}

//调用连接建立时钩子方法
func (s *Server) CallOnConnStartFunc(conn iface.Iconn) {
	if s.OnConnStartFunc != nil {
		Logger.Printf("[Server.CallOnConnStartFunc] connId=%d", conn.GetConnID())
		s.OnConnStartFunc(conn)
	}
}

//调用连接关闭时钩子方法
func (s *Server) CallOnConnStopFunc(conn iface.Iconn) {
	if s.OnConnStopFunc != nil {
		Logger.Printf("[Server.CallOnConnStopFunc] connId=%d", conn.GetConnID())
		s.OnConnStopFunc(conn)
	}
}
