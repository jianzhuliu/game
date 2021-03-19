package main

import (
	"runtime"

	"gitee.com/jianzhuliu/game/core"
	"gitee.com/jianzhuliu/game/examples/api"
	"gitee.com/jianzhuliu/game/examples/services"
)

func main() {
	//设置最大goroutine 调度器个数
	numGO := runtime.NumCPU() * 3 / 4
	if numGO < 1 {
		numGO = 1
	}
	runtime.GOMAXPROCS(numGO)

	//构建服务器
	s := core.NewServer("127.0.0.1", 8999)

	//添加路由
	api.Setup(s)

	//设置连接开启及关闭钩子方法
	s.SetOnConnStartFunc(services.OnConnStart)
	s.SetOnConnStopFunc(services.OnConnStop)

	//运行
	s.Run()
}
