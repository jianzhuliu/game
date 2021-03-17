package core

import (
	"sync"

	"gitee.com/jianzhuliu/game/iface"
)

//消息id路由处理对象
type Router struct {
	Handlers    map[uint32]iface.Ihandler //消息id对应处理函数
	handlerLock sync.RWMutex

	//开启工作池模式
	WorkerPoolSize int                   //工作池最多工人数
	TaskQueue      []chan iface.Irequest //每个工人的任务队列
}

//构造路由对象
func NewHandlers() *Router {
	return &Router{
		Handlers:       make(map[uint32]iface.Ihandler),
		WorkerPoolSize: ConfObj.WorkerPoolSize,
		TaskQueue:      make([]chan iface.Irequest, ConfObj.WorkerPoolSize),
	}
}

//添加消息处理函数
func (r *Router) AddRouter(msgId uint32, handler iface.Ihandler) {
	r.handlerLock.Lock()
	defer r.handlerLock.Unlock()

	r.Handlers[msgId] = handler
}

//执行处理
func (r *Router) DoHandle(request iface.Irequest) {
	r.handlerLock.RLock()
	defer r.handlerLock.RUnlock()
	if handler, ok := r.Handlers[request.GetMsgId()]; ok {
		//先后调用handler方法
		workerId := request.GetWorkerId()
		if workerId >= 0 {
			Logger.Printf("WorkerId(%d) handle connId=%d, msgId=%d", workerId, request.GetConnection().GetConnID(), request.GetMsgId())
		}

		handler.BeforeHandle(request)
		handler.HandleRequest(request)
		handler.AfterHandle(request)
	} else {
		Logger.Printf("[Router.DoHandle] =======> msgId(%d) has not defined handler", request.GetMsgId())
	}
}

///////////工作池模式

//开启服务
func (r *Router) OpenWorkerPool() {
	for i := 0; i < r.WorkerPoolSize; i++ {
		r.TaskQueue[i] = make(chan iface.Irequest, ConfObj.WorkerTaskSize)
		go func(i int, chanTask chan iface.Irequest) {
			Logger.Printf("[Router.OpenWorkerPool] worker(%d) start to work", i)
			for {
				select {
				case request := <-chanTask:
					go r.DoHandle(request)
				}
			}
		}(i, r.TaskQueue[i])
	}
}

//发送任务到工作池
func (r *Router) SendRequestToWorkerPool(request iface.Irequest) {
	connId := request.GetConnection().GetConnID()
	msgId := request.GetMsgId()
	workerId := (connId + int(msgId)) % ConfObj.WorkerPoolSize

	Logger.Printf("WorkerId(%d) get the job, connId=%d, msgId=%d", workerId, connId, msgId)

	//设置工作ID编号
	request.SetWorkerId(workerId)

	//随机分配到对应的工作池中
	r.TaskQueue[workerId] <- request
}
