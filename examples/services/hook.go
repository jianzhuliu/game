package services

import (
	"fmt"

	"gitee.com/jianzhuliu/game/core"
	"gitee.com/jianzhuliu/game/iface"
)

//连接开始时注册函数
func OnConnStart(conn iface.Iconn) {
	core.Logger.Print("[OnConnStart] connId=", conn.GetConnID())

	//设置属性连接id
	conn.WithValue("connId", conn.GetConnID())

	//创建一个玩家
	player := NewPlayer(uint32(conn.GetConnID()))

	//玩家绑定到连接
	conn.WithValue("player", player)

	//给其他连接发送上线消息
	chatObj := NewChatObj(fmt.Sprintf("player [%d] online", player.GetPlayerId()))
	chatObj.SendMsgToOthers(conn)
}

//连接关闭时注册函数
func OnConnStop(conn iface.Iconn) {
	core.Logger.Print("[OnConnStop] connId=", conn.GetConnID())

	//读取配置的属性
	connId := conn.Value("connId").(int)
	player := conn.Value("player").(*Player)
	core.Logger.Printf("get value, connId=%d, playerId=%d", connId, player.GetPlayerId())

	//给其他连接发送下线消息
	chatObj := NewChatObj(fmt.Sprintf("player [%d] logout", player.GetPlayerId()))
	chatObj.SendMsgToOthers(conn)
}
