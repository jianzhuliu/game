package iface

//消息id路由处理接口
type Ihandler interface {
	//处理请求前操作
	BeforeHandle(Irequest)

	//处理中
	HandleRequest(Irequest)

	//处理后操作
	AfterHandle(Irequest)
}
