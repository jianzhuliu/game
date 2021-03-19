package api

import (
	"gitee.com/jianzhuliu/game/examples/services"
	"gitee.com/jianzhuliu/game/iface"
)

func Setup(s iface.Iserver) {
	s.AddRouter(services.C_MSG_ID_CHAT, &ChatHandler{})
}
