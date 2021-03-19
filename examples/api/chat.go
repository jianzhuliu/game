package api

import (
	"fmt"

	"gitee.com/jianzhuliu/game/core"
	"gitee.com/jianzhuliu/game/iface"

	"gitee.com/jianzhuliu/game/examples/pb"
	"gitee.com/jianzhuliu/game/examples/services"

	"github.com/golang/protobuf/proto"
)

type ChatHandler struct {
	core.BaseHandler
}

func (c *ChatHandler) HandleRequest(request iface.Irequest) {
	//解析数据包
	msg := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		core.Logger.Print("[ChatHandler.HandleRequest] proto.Unmarshal err,", err)
		return
	}

	//获取当前玩家信息
	conn := request.GetConnection()
	player := conn.Value("player").(*services.Player)
	content := fmt.Sprintf("[%d]%s", player.GetPlayerId(), msg.GetContent())

	chatObj := services.NewChatObj(content)

	//给其他玩家广播
	chatObj.SendMsgToOthers(conn)
}
