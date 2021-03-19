# 项目产生来源

>-  学习开源项目 刘丹冰([@aceld](https://github.com/aceld/zinx)) ,自己也模仿着写了一个

>- 非常不错的项目，推荐初学者练手

# 知识点总结

## 面向接口编程

>- iface 目录下存放所有定义的接口

>- core 目录下为对应接口的实现封装对象

>- 多态, 比如，可自定义日志处理对象，只要实现了 iface.Ilogger 所有方法

>- 继承，比如, 提供了基本的 BaseHandler, 自定义的处理对象只需嵌套它即可

## 连接读写分离

>- 连接建立，开启读写分离，通过管道ChanMsg chan iface.Imsg 进行数据沟通

## 连接管理器

>- ConnMgr 管理所有连接，方便最大连接数限制及服务器退出时统一关闭所有连接

## 数据协议采用 LTV(长度+消息类型+数据值) 形式

>- 数据的打包与解包，统一由 Dbpack 处理

## 消息队列工作池模式

>- 开启固定个数工人干活，每个工人负责处理各自的消息队列

>- 避免每个消息开启一个 goroutine，减少了 goroutine 数量

## 路由器模式

>- 根据消息类型，定义不同的消息路由

## 连接对象可设置各自属性

>- 采用 context.Context 接口,配合 context.WithValue 处理

## 支持配置连接建立及退出钩子方法

>- 设置连接建立时钩子方法 SetOnConnStartFunc(func(Iconn))
>- 设置连接关闭时钩子方法 SetOnConnStopFunc(func(Iconn))

# 简单 demo 

```go 
package main

import (
	"runtime"

	"gitee.com/jianzhuliu/game/core"
	"gitee.com/jianzhuliu/game/iface"
)

const (
	C_MSG_ID_HELLO = 1
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

//连接开始时注册函数
func OnConnStart(conn iface.Iconn) {
	core.Logger.Print("[OnConnStart] connId=", conn.GetConnID())

	//设置属性
	conn.WithValue("name", "jianzhu")
	conn.WithValue("gender", 1)
}

//连接关闭时注册函数
func OnConnStop(conn iface.Iconn) {
	core.Logger.Print("[OnConnStop] connId=", conn.GetConnID())

	//读取配置的属性
	name := conn.Value("name").(string)
	gender := conn.Value("gender").(int)
	core.Logger.Printf("get value, name=%s, gender=%d", name, gender)

}

func main() {
	s := core.NewServer("127.0.0.1", 8999)
	s.AddRouter(C_MSG_ID_HELLO, &HelloHandler{})

	s.SetOnConnStartFunc(OnConnStart)
	s.SetOnConnStopFunc(OnConnStop)

	s.Run()
}


```
