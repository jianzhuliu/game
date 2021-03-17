package core

import "gitee.com/jianzhuliu/game/iface"

//定义消息处理基类
type BaseHandler struct {
}

//处理请求前操作
func (h *BaseHandler) BeforeHandle(iface.Irequest) {

}

//处理中
func (h *BaseHandler) HandleRequest(iface.Irequest) {

}

//处理后操作
func (h *BaseHandler) AfterHandle(iface.Irequest) {

}
