package services

import (
	"gitee.com/jianzhuliu/game/core"
	"gitee.com/jianzhuliu/game/iface"

	"gitee.com/jianzhuliu/game/examples/pb"
	"github.com/golang/protobuf/proto"
)

//聊天对象
type ChatObj struct {
	Content string
}

//创建聊天对象
func NewChatObj(content string) *ChatObj {
	return &ChatObj{
		Content: content,
	}
}

//给在线玩家发送消息
func (c *ChatObj) SendMsgToOthers(conn iface.Iconn) {
	//构造数据包
	proto_data := &pb.Talk{
		Content: c.Content,
	}

	//数据打包
	data, err := proto.Marshal(proto_data)
	if err != nil {
		core.Logger.Print("[ChatObj.SendMsgToAll] proto.Marshal err,", err)
		return
	}

	conns := conn.GetServer().GetConnMgr().GetAllConns()
	connId := conn.GetConnID()

	if len(conns) > 1 {
		for _, connection := range conns {
			if connection.GetConnID() != connId {
				connection.SendMsg(C_MSG_ID_CHAT, data)
			}
		}
	}
}
