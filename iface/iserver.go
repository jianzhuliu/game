package iface

//服务器抽象接口
type Iserver interface {
	Run()  //运行
	Stop() //停止服务

	GetRouter() Irouter         //获取路由处理对象
	AddRouter(uint32, Ihandler) //添加消息路由
	GetConnMgr() Iconnmgr       //获取连接管理器对象
	SetLogger(Ilogger)          //设置日志对象

	//设置连接建立时钩子方法
	SetOnConnStartFunc(func(Iconn))

	//设置连接关闭时钩子方法
	SetOnConnStopFunc(func(Iconn))

	//调用连接建立时钩子方法
	CallOnConnStartFunc(Iconn)

	//调用连接关闭时钩子方法
	CallOnConnStopFunc(Iconn)
}
