package iface

//消息id路由处理接口
type Irouter interface {
	//添加消息处理函数
	AddRouter(uint32, Ihandler)

	//执行处理
	DoHandle(Irequest)

	//开启工作池
	OpenWorkerPool()

	//发送请求到工作池
	SendRequestToWorkerPool(Irequest)
}
