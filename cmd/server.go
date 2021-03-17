package main

import (
	"runtime"

	"gitee.com/jianzhuliu/game/core"
	"gitee.com/jianzhuliu/game/iface"
)

const (
	C_MSG_ID_HELLO = 1
	C_MSG_ID_HOME  = 2
)

//模拟服务器端
type HelloHandler struct {
	core.BaseHandler
}

func (h *HelloHandler) HandleRequest(request iface.Irequest) {
	//源数据放回
	core.Logger.Printf("[HelloHandler] =====> msgId=%d, msgData=%v \n", request.GetMsgId(), string(request.GetData()))
	request.SendMsg(C_MSG_ID_HELLO, request.GetData())
}

type HomeHandler struct {
	core.BaseHandler
}

func (h *HomeHandler) HandleRequest(request iface.Irequest) {
	//源数据放回
	core.Logger.Printf("[HelloHandler] =====> msgId=%d, msgData=%v \n", request.GetMsgId(), string(request.GetData()))
	request.SendMsg(C_MSG_ID_HOME, request.GetData())
}

func main() {
	//不全部占用cpu个数
	n := runtime.NumCPU() * 3 / 4
	if n < 1 {
		n = 1
	}

	runtime.GOMAXPROCS(n)

	s := core.NewServer("127.0.0.1", 8999)
	s.AddRouter(C_MSG_ID_HELLO, &HelloHandler{})
	s.AddRouter(C_MSG_ID_HOME, &HomeHandler{})

	s.Run()
}
